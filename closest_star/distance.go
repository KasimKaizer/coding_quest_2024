package closeststar

import (
	"math"
)

type SolarData struct {
	Name string
	X    float64
	Y    float64
	Z    float64
}

func Calculate(data []*SolarData, smallest float64) float64 {
	dataLen := len(data)
	for i := 0; i < dataLen-1; i++ {
		curData := data[i]
		for _, item := range data[i+1:] {
			num := distance(curData, item)
			if num < smallest {
				smallest = num
			}
		}
	}
	return smallest
}

func distance(first, second *SolarData) float64 {
	sum := math.Pow(first.X-second.X, 2) + math.Pow(first.Y-second.Y, 2) + math.Pow(first.Z-second.Z, 2)
	return math.Sqrt(sum)
}
