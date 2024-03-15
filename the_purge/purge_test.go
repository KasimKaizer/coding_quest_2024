package purge

import "testing"

var TestCases = []struct {
	Description string
	InputFile   string
	Expected    int64
}{
	{
		"Base Test Case",
		"base_test.txt",
		103879262,
	},
	{
		"Real Test Case",
		"day07_files_to_scan.txt",
		349035592144,
	},
}

func TestDelete(t *testing.T) {
	for _, test := range TestCases {
		t.Run(test.Description, func(t *testing.T) {
			got, err := Delete(test.InputFile)
			if err != nil {
				t.Fatalf("got error: %v", err)
			}
			if got != test.Expected {
				t.Fatalf("expected: %d, got: %d", test.Expected, got)
			}
		})
	}
}
