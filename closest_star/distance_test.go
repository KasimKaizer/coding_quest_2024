package closeststar

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"testing"
)

// parse the table, we can split by the at least

var CalTestCases = []struct {
	Description string
	FilePath    string
	Expected    float64
}{
	{
		"Base_Case_Test",
		"test_1.txt",
		3.508,
	},
	{
		"real_Case_Test",
		"test_2.txt",
		1.099,
	},
}

var SplitTestCases = []struct {
	Description string
	Data        string
	Expected    []string
}{
	{
		"Base_Case_Test",
		"WISE J085510.83-071442.5               7.532    -3.967    -5.664     2.985",
		[]string{
			"WISE J085510.83-071442.5",
			"7.532", "-3.967", "-5.664", "2.985",
		},
	},
}

func TestCalculate(t *testing.T) {
	for _, test := range CalTestCases {
		t.Run(test.Description, func(t *testing.T) {
			data, small := constructData(test.FilePath)
			temp := math.Round(Calculate(data, small) * 1000)
			got := int(temp)
			if got != int(test.Expected*1000) {
				log.Fatal("not working")
			}
		})
	}
}

func BenchmarkCalculate(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	data, small := constructData(CalTestCases[1].FilePath)
	for i := 0; i < b.N; i++ {
		Calculate(data, small)
	}
}

func Split(data string) []string {
	dataLen, start := len(data), 0
	out := make([]string, 0, 5)

	for i := 0; i <= dataLen; i++ {
		if i == dataLen-1 || i != dataLen && !(data[i] == ' ' && data[i+1] == ' ') {
			continue
		}
		out = append(out, data[start:i])

		for i < dataLen && data[i] == ' ' {
			i++
		}
		start = i
	}
	return out
}

func Sanitize(data []string) (*SolarData, float64) {
	out := SolarData{
		Name: data[0],
	}

	var err error
	out.X, err = strconv.ParseFloat(data[2], 64)
	if err != nil {
		log.Fatal("closeststar.Sanitize: error while parsing float for X")
	}
	out.Y, err = strconv.ParseFloat(data[3], 64)
	if err != nil {
		log.Fatal("closeststar.Sanitize: error while parsing float for y")
	}
	out.Z, err = strconv.ParseFloat(data[4], 64)
	if err != nil {
		log.Fatal("closeststar.Sanitize: error while parsing float for y")
	}

	distance, err := strconv.ParseFloat(data[1], 64)
	if err != nil {
		log.Fatal("closeststar.Sanitize: error while parsing float for distance")
	}

	return &out, distance
}

func constructData(filePath string) ([]*SolarData, float64) {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	firstLine := true
	scanner := bufio.NewScanner(f)
	small := math.MaxFloat64

	out := make([]*SolarData, 0)
	for scanner.Scan() {
		if firstLine {
			firstLine = false
			continue
		}
		curData, num := Sanitize(Split(scanner.Text()))
		if num < small {
			small = num
		}
		out = append(out, curData)
	}
	return out, small
}
