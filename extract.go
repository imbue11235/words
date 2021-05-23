package words

import (
	"unicode"
	"unicode/utf8"
)

const (
	symbol = 1 + iota
	uppercase
	lowercase
	space
	digit
	punctuation
	unknown
)

const hyphen = rune(45)

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
	case unicode.IsPunct(r):
		return punctuation
	}

	return unknown
}

func shouldInclude(runeType int, config *config) bool {
	switch runeType {
	case symbol:
		return config.includeSymbols
	case punctuation:
		return config.includePunctuation
	case space:
		return config.includeSpaces
	}

	return true
}

// Rules:
// 1. Invalid UTF8-strings will not be split

func extract(input string, config *config) []string {
	if !utf8.ValidString(input) {
		return []string{input}
	}

	var runes [][]rune
	var runeType int
	lastRuneType := 0
	indexes := -1

	for _, r := range input {
		runeType = getRuneType(r)

		if config.allowHyphenatedWords && r == hyphen {
			runes[len(runes) - 1] = append(runes[len(runes) - 1], r)
			continue
		}

		if !shouldInclude(runeType, config) {
			lastRuneType = runeType
			continue
		}

		if runeType == lastRuneType {
			runes[len(runes) - 1] = append(runes[len(runes) - 1], r)
			lastRuneType = runeType
			continue
		}

		runes = append(runes, []rune{r})
		indexes++

		if lastRuneType == uppercase && runeType == lowercase {
			// Prepend the last character of previous rune-slice
			runes[indexes] = append([]rune{runes[indexes-1][len(runes[indexes-1])-1]}, runes[indexes]...)
			// Remove the last character from the previous rune-slice
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

func Extract(input string, options ...Option) []string {
	config := newDefaultConfig()
	config.apply(options...)

	return extract(input, config)
}