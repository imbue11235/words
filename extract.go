// Package words provides capabilities for splitting
// a string into a slice of words by a collection of rules
//
// Rules:
// 1. Invalid UTF8-strings will not be split
// 2. Hyphenated words will be treated as individual words unless disabled. E.g. "small-town" => []{"small", "town"}
// 3. If the character is a space, punctuation or symbol, it will be voided,
// unless disabled. E.g. "my_string  here" => []{"my", "string", "here"}
// 4. Characters of same type in sequence, will be put together.
// 5. If the current character is a lowercase, and the last character of the previous word was uppercase,
// the uppercase letter will be moved to the lowercase string. E.g. "YAMLParser" => []{"YAML", "Parser"}
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

// getRuneKind takes the rune and returns
// the int representation of it's kind
func getRuneKind(r rune) int {
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

// shouldInclude checks if the kind of rune should be included
// in the word
func shouldInclude(runeKind int, config *config) bool {
	switch runeKind {
	case symbol:
		return config.includeSymbols
	case punctuation:
		return config.includePunctuation
	case space:
		return config.includeSpaces
	}

	return true
}

func in(runeKind int, runeKinds []int) bool {
	for _, kind := range runeKinds {
		if runeKind == kind {
			return true
		}
	}

	return false
}

func isHyphenatedWord(r rune, lastRuneKind, nextRuneKind int) bool {
	if r != hyphen {
		return false
	}

	// Make sure that the runes kind are equal
	// to avoid false hyphenated words.
	// E.g. "SOME-word", should still be []{"SOME", "word"}
	if lastRuneKind != nextRuneKind {
		return false
	}

	return in(lastRuneKind, []int{lowercase, uppercase}) && in(nextRuneKind, []int{lowercase, uppercase})
}

// extract with by the defined rules
func extract(input string, config *config) []string {
	// Early return, if invalid string (Rule 1)
	if !utf8.ValidString(input) {
		return []string{input}
	}

	var runes [][]rune
	runeKind, lastRuneKind, runesLen := 0, 0, -1

	for i, r := range input {
		// If hyphenated words are allowed and current character is hyphenated,
		// it'll get appended to the current rune slice,
		// if the adjacent runes of a hyphen is a letter of same kind (upper/lowercase),
		// without keeping track of it's rune type (Rule 2).
		if config.allowHyphenatedWords {
			var nextRuneKind int

			if len(input) > i+1 {
				nextRuneKind = getRuneKind(rune(input[i+1]))
			}

			if isHyphenatedWord(r, lastRuneKind, nextRuneKind) {
				runes[runesLen] = append(runes[runesLen], r)
				continue
			}
		}

		// Define the rune kind
		runeKind = getRuneKind(r)

		// Determine if the current rune should be voided or not.
		// The current rune kind will still be set, to make sure that a new word
		// will be started on in next iteration (Rule 3).
		if !shouldInclude(runeKind, config) {
			lastRuneKind = runeKind
			continue
		}

		// If the rune has same kind as last rune, it will get appended
		// to the current word. (Rule 4)
		if runeKind == lastRuneKind {
			runes[runesLen] = append(runes[runesLen], r)
			continue
		}

		// Start a new word
		runes = append(runes, []rune{r})

		// Keep track of the runes index, instead of using len(runes) to find current index
		runesLen++

		// Move a uppercase rune from the end of previous word, to this word (Rule 5).
		if lastRuneKind == uppercase && runeKind == lowercase {
			// Prepend the last character of previous rune-slice
			runes[runesLen] = append([]rune{runes[runesLen-1][len(runes[runesLen-1])-1]}, runes[runesLen]...)
			// Remove the last character from the previous rune-slice
			runes[runesLen-1] = runes[runesLen-1][:len(runes[runesLen-1])-1]
		}

		lastRuneKind = runeKind

	}

	// Convert the rune slices to strings
	var output []string
	for _, r := range runes {
		if len(r) == 0 {
			continue
		}

		output = append(output, string(r))
	}

	return output
}

// Extract extracts words from a given string with potential options.
func Extract(input string, options ...Option) []string {
	config := newDefaultConfig()
	config.apply(options...)

	return extract(input, config)
}