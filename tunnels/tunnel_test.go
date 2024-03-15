package tunnels

import (
	"bufio"
	"bytes"
	"os"
	"testing"
)

// first create something which parses our text into a 3d matrix
// we can just read from the file, if the line begins with a empty space then we have found
// next layer

var TestCases = []struct {
	Description string
	MazeFile    string
	Start       Position
	Expected    int
}{
	{
		"Base Case Test",
		"base_case.txt",
		Position{0, 1, 0},
		52,
	},
	{
		"Real Case Test",
		"real_case.txt",
		Position{0, 1, 0},
		250,
	},
}

func createMaze(filePath string) ([][][]byte, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	maze := make([][][]byte, 2)
	curLayer := 0
	for scanner.Scan() {
		text := bytes.Clone(scanner.Bytes())
		if len(text) == 0 || text[1] == ' ' {
			curLayer = 1 - curLayer
			continue
		}
		maze[curLayer] = append(maze[curLayer], text)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return maze, nil
}

func TestFindShortestPath(t *testing.T) {
	for _, test := range TestCases {
		t.Run(test.Description, func(t *testing.T) {
			maze, err := createMaze(test.MazeFile)
			if err != nil {
				t.Fatal("encountered err while trying to generate the maze, tests are broken")
			}
			got := FindShortestPath(maze, test.Start)
			if got != test.Expected {
				t.Fatalf("expected: %d, got: %d", test.Expected, got)
			}
		})
	}

}

func BenchmarkFindShortestPath(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode")
	}
	for i := 0; i < b.N; i++ {
		for _, test := range TestCases {
			maze, _ := createMaze(test.MazeFile)
			FindShortestPath(maze, test.Start)
		}
	}

}
