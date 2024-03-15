package key

import "strings"

func Decode(input []int) string {
	var out strings.Builder
	count := 0
	for i := 0; i < len(input); i++ {
		if i%2 == 0 {
			repeat(&out, &count, input[i], '.')
			continue
		}
		repeat(&out, &count, input[i], '#')
	}
	return out.String()
}

func repeat(res *strings.Builder, count *int, len int, sym byte) {
	for i := 0; i < len; i++ {
		res.WriteByte(sym)
		*count++
		if *count == 100 {
			res.WriteByte('\n')
			*count = 0
		}
	}
}
