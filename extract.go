package words

import (
	"unicode"
	"unicode/utf8"
)

const (
	symbol = 1 << iota
	uppercase
	lowercase
	space
	digit
	unknown
)

func getRuneType(r rune) int {
	switch {
	case unicode.IsSymbol(r):
		return symbol
	case unicode.IsUpper(r):
		return uppercase
	case unicode.IsLower(r):
		return lowercase
	case unicode.IsSpace(r):
		return space
	case unicode.IsDigit(r):
		return digit
	}

	return unknown
}


// Rules:
// 1. Invalid UTF8-strings will not be split

func Extract(input string) []string {
	if !utf8.ValidString(input) {
		return []string{input}
	}

	var runes [][]rune
	var runeType int
	lastRuneType := 0
	indexes := -1

	for _, r := range input {
		runeType = getRuneType(r)

		if runeType == lastRuneType {
			runes[len(runes) - 1] = append(runes[len(runes) - 1], r)
			lastRuneType = runeType
			continue
		}

		runes = append(runes, []rune{r})
		indexes++

		if lastRuneType == uppercase && runeType == lowercase {
			runes[indexes] = append([]rune{runes[indexes-1][len(runes[indexes-1])-1]}, runes[indexes]...)
			runes[indexes-1] = runes[indexes-1][:len(runes[indexes-1])-1]
		}

		lastRuneType = runeType
	}

	var output []string
	for _, r := range runes {
		if len(r) == 0 {
			continue
		}

		output = append(output, string(r))
	}

	return output
}