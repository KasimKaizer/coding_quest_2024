package firewall

import (
	"log"
	"testing"
)

var InspectTestCases = []struct {
	name     string
	input    string
	expected string
}{
	{
		"First base test case",
		"test1.txt",
		"868/1625",
	},
	{
		"Main test case",
		"test2.txt",
		"258956/256237",
	},
}

var NewIpTestCases = []struct {
	name     string
	input    string
	expected IpAddress
}{
	{
		"First hex ip test case",
		"0A000bc1d7253441",
		IpAddress{"10.0.11.193", "215.37.52.65"},
	},
	{
		"Second hex ip test case",
		"c0a800b833c555ee",
		IpAddress{"192.168.0.184", "51.197.85.238"},
	},
}

func TestInspect(t *testing.T) {
	for _, test := range InspectTestCases {
		t.Run(test.name, func(t *testing.T) {
			got, err := Inspect(test.input)
			if err != nil {
				log.Fatalf("got error: %v", err)
			}
			if got != test.expected {
				log.Fatalf("Expected: %s, Got: %s", test.expected, got)
			}
		})
	}
}

func TestNewIp(t *testing.T) {
	for _, test := range NewIpTestCases {
		t.Run(test.name, func(t *testing.T) {
			got, err := NewIp(test.input)
			if err != nil {
				log.Fatalf("got error: %v", err)
			}
			if *got != test.expected {
				if got.dest != test.expected.dest {
					log.Fatalf("destination not equal. Expected: %s, Got: %s", test.expected, got)
				}
				if got.sour != test.expected.sour {
					log.Fatalf("source not equal. Expected: %s, Got: %s", test.expected, got)
				}

			}
		})
	}
}
