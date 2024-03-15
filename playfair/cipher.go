package playfair

import (
	"strings"
)

// add 'x' to the end if the number of characters in the cipher are odd.
// if the last character is 'x' then then add 'q'.
// replace all occurrence of 'j' with 'i'.
// strip all spaces.
//
// the key should not have any repeating characters.
//
// create a function to sanitize the key.
// create a function to sanitize the input text.
// create a 5 x 5 grid with 25 characters.
//
// this cipher won't work with utf-8 characters so there is no point accounting for them, its
// also gonna highly effect the performance of this solution.

const (
	decrypt = 4
	encrypt = 1
)

type grid [][]byte

func Encrypt(key, text string) string {
	return cipher(key, text, encrypt)
}
func Decrypt(key, text string) string {
	return cipher(key, text, decrypt)

}

func cipher(key, text string, ops int) string {
	newKey := sanitizeKey(key)
	newText := sanitizeText(text)
	grid := newGrid(newKey)
	oldIdx := 0
	var secret strings.Builder
	for i := 0; i < len(newText)-1; i = i + 2 {
		var res [2]byte
		rowPos1, colPos1 := grid.position(newText[i])
		rowPos2, colPos2 := grid.position(newText[i+1])
		switch {
		case colPos1 == colPos2:
			res[0] = grid[(rowPos1+ops)%5][colPos1]
			res[1] = grid[(rowPos2+ops)%5][colPos1]
		case rowPos1 == rowPos2:
			res[0] = grid[rowPos1][(colPos1+ops)%5]
			res[1] = grid[rowPos2][(colPos2+ops)%5]
		default:
			res[0] = grid[rowPos1][colPos2]
			res[1] = grid[rowPos2][colPos1]
		}
		for _, char := range res {
			if text[oldIdx] == ' ' {
				secret.WriteByte(' ')
				oldIdx++
			}
			secret.WriteByte(char)
			oldIdx++
		}
	}
	return secret.String()
}

func sanitizeText(text string) string {
	var newText strings.Builder
	lastX := false
	for i := 0; i < len(text); i++ {
		if text[i] == ' ' {
			continue
		}
		if text[i] == 'j' {
			newText.WriteByte('i')
			continue
		}
		if i == len(text)-1 && text[i] == 'x' {
			lastX = true
		}
		newText.WriteByte(text[i])
	}
	if newText.Len()%2 != 0 {
		if lastX {
			newText.WriteByte('q')
		} else {
			newText.WriteByte('x')
		}
	}
	return newText.String()
}

func sanitizeKey(key string) string {
	mem := make(map[byte]struct{})
	var newKey strings.Builder
	for i := 0; i < len(key); i++ {
		if _, ok := mem[key[i]]; ok {
			continue
		}
		newKey.WriteByte(key[i])
		mem[key[i]] = struct{}{}
	}
	return newKey.String()
}

func newGrid(key string) grid {
	grid := make(grid, 5)
	mem := make(map[byte]bool)
	mem['j'] = true
	keyIdx, sLetter := 0, byte('a')
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if keyIdx < len(key) {
				grid[i] = append(grid[i], key[keyIdx])
				mem[key[keyIdx]] = true
				keyIdx++
				continue
			}
			for mem[sLetter] {
				sLetter++
			}
			grid[i] = append(grid[i], sLetter)
			sLetter++
		}
	}
	return grid
}

func (g grid) position(x byte) (int, int) {
	for row, line := range g {
		for col, char := range line {
			if char == x {
				return row, col
			}
		}
	}
	return -1, -1
}
