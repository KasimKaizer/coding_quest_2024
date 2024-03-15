package rover

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

type RoverData [][]int
type Outposts map[string]int

var errUnknownOutpost = errors.New("rover: unknown outpost in the provided path")

func ParseData(filePath string) (Outposts, RoverData, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var notFirst bool
	data, Outposts := make(RoverData, 0), make(Outposts)
	for scanner.Scan() {
		if !notFirst {
			splitData := strings.Fields(scanner.Text())
			for i := 0; i < len(splitData); i++ {
				Outposts[splitData[i]] = i
			}
			notFirst = true
			continue
		}
		splitData := strings.Fields(scanner.Text())
		dist := make([]int, len(splitData)-1)
		for idx, numStr := range splitData[1:] {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				return nil, nil, err
			}
			dist[idx] = num
		}
		data = append(data, dist)
	}
	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}
	return Outposts, data, nil
}

func CalMultiDist(rovers []string, outposts Outposts, data RoverData) (int, error) {
	total := 0
	for _, rover := range rovers {
		dist, err := calDistance(rover, outposts, data)
		if err != nil {
			return 0, err
		}
		total += dist
	}
	return total, nil
}

func calDistance(path string, outposts Outposts, data RoverData) (int, error) {
	proper := strings.Split(path, ": ")
	sanPath := strings.Split(proper[1], " -> ")
	total := 0
	for i := 0; i < len(sanPath)-1; i++ {
		from, ok := outposts[sanPath[i]]
		if !ok {
			return 0, errUnknownOutpost
		}
		to, ok := outposts[sanPath[i+1]]
		if !ok {
			return 0, errUnknownOutpost
		}
		total += data[from][to]
	}
	return total, nil
}
