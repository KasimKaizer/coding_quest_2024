package codingquest

import (
	"bufio"
	"errors"
	"math"
	"os"
	"strconv"
	"strings"
)

func Purchase(file string) (int, error) {
	f, err := os.Open(file)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	flights := make(map[string]int)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		parts := strings.Fields(scanner.Text())
		if len(parts) < 3 {
			return 0, errors.New("invalid text")
		}
		num, err := strconv.Atoi(parts[2])
		if err != nil {
			return 0, err
		}
		switch parts[1] {
		case "Seat", "Luggage", "Tax", "Meals", "Fee":
			flights[parts[0]] += num
		case "Discount", "Rebate":
			flights[parts[0]] -= num
		}
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}

	lowestPrice := math.MaxInt

	for _, price := range flights {
		if price > lowestPrice {
			continue
		}
		lowestPrice = price
	}
	return lowestPrice, nil
}
