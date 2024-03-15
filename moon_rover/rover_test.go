package rover

import (
	"io"
	"os"
	"reflect"
	"strings"
	"testing"
)

var ParseDataTestCases = []struct {
	Description string
	Input       string
	Outposts    Outposts
	Data        RoverData
}{
	{
		"Base Test Case",
		"test_1.txt",
		Outposts{"base": 0, "ta00": 1, "cx22": 2, "xj84": 3},
		RoverData{{0, 55457, 63529, 61302}, {55457, 0, 111890, 35768}, {63529, 111890, 0, 98977}, {61302, 35768, 98977, 0}},
	},
}

var CalDistanceTestCases = []struct {
	Description string
	Paths       string
	dataFile    string
	Expected    int
}{
	{
		"Base Test Case",
		"pathtest_1.txt",
		"test_1.txt",
		353480,
	},
	{
		"Real Test Case",
		"pathtest_2.txt",
		"test_2.txt",
		6979629,
	},
}

func TestParseData(t *testing.T) {
	for _, test := range ParseDataTestCases {
		t.Run(test.Description, func(t *testing.T) {
			rover, data, err := ParseData(test.Input)
			if err != nil {
				t.Fatal("Got unexpected error")
			}
			if !reflect.DeepEqual(rover, test.Outposts) {
				t.Fatalf("Outposts aren't right. want: %+v, got: %+v", test.Outposts, rover)
			}
			if !reflect.DeepEqual(test.Data, data) {
				t.Fatalf("out data is incorrect. want: %+v, got: %+v", test.Data, data)
			}
		})
	}
}

func TestCalMultiDist(t *testing.T) {
	for _, test := range CalDistanceTestCases {
		t.Run(test.Description, func(t *testing.T) {
			outpost, data, err := ParseData(test.dataFile)
			if err != nil {
				t.Fatal(err)
			}
			f, err := os.Open(test.Paths)
			if err != nil {
				t.Fatal(err)
			}
			defer f.Close()
			raw, err := io.ReadAll(f)
			if err != nil {
				t.Fatal(err)
			}
			paths := strings.Split(string(raw), "\n")
			got, err := CalMultiDist(paths, outpost, data)
			if err != nil {
				t.Fatal(err)
			}
			if got != test.Expected {
				t.Fatalf("wanted: %d, got: %d", test.Expected, got)
			}
		})
	}
}
