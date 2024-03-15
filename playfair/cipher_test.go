package playfair

import (
	"testing"
)

var testCasesDecrypt = []struct {
	description, key, input, expected string
}{
	{"Base test case",
		"helloworld",
		"wp nehslv ewgw",
		"et phonex home",
	},
	{
		"Real case",
		"codingquest",
		"rmqfgs yegv em qnpu pdml dc atuy olzy anpu",
		"please pick up some milk on thex wayx home",
	},
}

var testCasesEncrypt = []struct {
	description, key, input, expected string
}{
	{"Base test case",
		"helloworld",
		"et phonex home",
		"wp nehslv ewgw",
	},
	{
		"Real case",
		"codingquest",
		"please pick up some milk on thex wayx home",
		"rmqfgs yegv em qnpu pdml dc atuy olzy anpu",
	},
}

func TestEncrypt(t *testing.T) {
	for _, test := range testCasesEncrypt {
		t.Run(test.description, func(t *testing.T) {
			got := Encrypt(test.key, test.input)
			if got != test.expected {
				t.Fatalf("expected: %s, got: %s", test.expected, got)
			}
		})
	}
}

func TestDecrypt(t *testing.T) {
	for _, test := range testCasesDecrypt {
		t.Run(test.description, func(t *testing.T) {
			got := Decrypt(test.key, test.input)
			if got != test.expected {
				t.Fatalf("expected: %s, got: %s", test.expected, got)
			}
		})
	}
}

func BenchmarkDecrypt(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode")
	}
	for i := 0; i < b.N; i++ {
		for _, test := range testCasesDecrypt {
			Decrypt(test.key, test.input)
		}
	}
}
