package connect

import "testing"

var TestCases = []struct {
	Description string
	Steps       []int
	Max         int
	Expected    string
}{
	{
		"Base Test Case",
		[]int{3, 2, 1},
		5,
		"13",
	},
	{
		"Real Test Case",
		[]int{40, 12, 2, 1},
		856,
		"361595632332583638761407421958897298379960745882500550853575978681928496636233758054533916390012124244431806190608039087381666468880612638124124662565287224989590899000769252066051",
	},
}

func TestFindPath(t *testing.T) {
	for _, test := range TestCases {
		t.Run(test.Description, func(t *testing.T) {
			got := FindPath(test.Steps, test.Max)
			if got.String() != test.Expected {
				t.Fatalf("expected: %s, got: %s", test.Expected, got)
			}
		})
	}

}
