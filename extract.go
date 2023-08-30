// Package words provides capabilities for splitting
// a string into a slice of words by a collection of rules
package words

import (
	"golang.org/x/exp/slices"
	"unicode"
	"unicode/utf8"
)

// Rules:
// 1. Invalid UTF8-strings will not be split
// 2. Hyphenated words will be treated as individual words unless disabled. E.g. "small-town" => []{"small", "town"}
// 3. If the character is a space, punctuation or symbol, it will be voided,
// unless disabled. E.g. "my_string  here" => []{"my", "string", "here"}
// 4. Characters of same type in sequence, will be put together.
// 5. If the current character is a lowercase, and the last character of the previous word was uppercase,
// the uppercase letter will be moved to the lowercase string. E.g. "YAMLParser" => []{"YAML", "Parser"}

type RuneKind int

const (
	Symbol RuneKind = 1 + iota
	Uppercase
	Lowercase
	Space
	Digit
	Punctuation
	Unknown
)

const hyphen = rune(45)

// getRuneKind takes the rune and returns
// the int representation of it's kind
func getRuneKind(r rune) RuneKind {
	switch {
	case unicode.IsSymbol(r):
		return Symbol
	case unicode.IsUpper(r):
		return Uppercase
	case unicode.IsLower(r):
		return Lowercase
	case unicode.IsSpace(r):
		return Space
	case unicode.IsDigit(r):
		return Digit
	case unicode.IsPunct(r):
		return Punctuation
	}

	return Unknown
}

// shouldInclude checks if the kind of rune should be included
// in the word
func shouldInclude(runeKind RuneKind, config *config) bool {
	switch runeKind {
	case Symbol:
		return config.includeSymbols
	case Punctuation:
		return config.includePunctuation
	case Space:
		return config.includeSpaces
	}

	return true
}

// isHyphenatedWord determines if the word is a hyphenated word
// by looking at adjacent rune kinds
func isHyphenatedWord(r rune, lastRuneKind, nextRuneKind RuneKind) bool {
	if r != hyphen {
		return false
	}

	// Make sure that the runes kind are equal
	// to avoid false hyphenated words.
	// E.g. "SOME-word", should still be []{"SOME", "word"}
	if lastRuneKind != nextRuneKind {
		return false
	}

	return slices.Contains([]RuneKind{Lowercase, Uppercase}, lastRuneKind) && slices.Contains([]RuneKind{Lowercase, Uppercase}, nextRuneKind)
}

// extract with by the defined rules
func extract(input string, config *config) []string {
	// Early return, if invalid string (Rule 1)
	if !utf8.ValidString(input) {
		return []string{input}
	}

	var runes [][]rune
	var runeKind, lastRuneKind RuneKind
	runesLen := -1

	for i, r := range input {
		// If the rune should be ignored, we will simply add it to
		// the current word, and treat it of same rune kind as the last
		// added value
		if slices.Contains(config.ignoredRunes, r) || slices.Contains(config.ignoredRunesKinds, getRuneKind(r)) {
			// If the current rune is the first rune, we will append a new slice
			if runesLen == -1 {
				runes = append(runes, []rune{})
				runesLen++
			}

			runes[runesLen] = append(runes[runesLen], r)

			// If there is a next rune, we will set the last rune kind to the next rune kind
			// to indicate that the next rune should be treated as the same kind as the last,
			// even if it's not.
			if len(input) > i+1 {
				lastRuneKind = getRuneKind(rune(input[i+1]))
			}

			continue
		}

		// If hyphenated words are allowed and current character is hyphenated,
		// it'll get appended to the current rune slice,
		// if the adjacent runes of a hyphen is a letter of same kind (upper/lowercase),
		// without keeping track of it's rune type (Rule 2).
		if config.allowHyphenatedWords {
			var nextRuneKind RuneKind

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

		// Move an uppercase rune from the end of previous word, to this word (Rule 5).
		if lastRuneKind == Uppercase && runeKind == Lowercase {
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
	cfg := newDefaultConfig()
	cfg.apply(options...)

	return extract(input, cfg)
}
