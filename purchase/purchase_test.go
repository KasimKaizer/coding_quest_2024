package codingquest

import (
	"fmt"
	"log"
	"testing"
)

var tests = []struct {
	name     string
	input    string
	expected int
}{
	{
		"First case",
		"test.txt",
		2533,
	},
	{
		"Second case",
		"test2.txt",
		0,
	},
}

func TestPurchase(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res, err := Purchase(test.input)
			if err != nil {
				log.Fatal(err)
			}
			if test.input == "First case" {
				if test.expected != res {
					log.Fatalf("This function doesn't work, expected: %d, got: %d", test.expected, res)
				}
				return
			}
			fmt.Print(res)
		})
	}
}
