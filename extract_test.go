package words_test

import (
	"github.com/imbue11235/words"
	"reflect"
	"testing"
)

type testSet struct {
	input string
	expected []string
}

func runExtractTest(t *testing.T, tests []testSet, options ...words.Option) {
	for _, test := range tests {
		extraction := words.Extract(test.input, options...)

		// If both slices are empty, just continue
		if len(extraction) == 0 && len(test.expected) == 0 {
			continue
		}

		if !reflect.DeepEqual(extraction, test.expected) {
			t.Errorf("Expected %v to be %v", extraction, test.expected)
		}
	}
}

func TestExtract(t *testing.T) {
	tests := []testSet{
		{"", []string{}},
		{"100cm", []string{"100", "cm"}},
		{"μου αρέσουν τα μπιφτέκια", []string{"μου", "αρέσουν", "τα", "μπιφτέκια"}},
		{"aeiouAreVowels", []string{"aeiou", "Are", "Vowels"}},
		{"XmlHTTP", []string{"Xml", "HTTP"}},
		{"isISO8601", []string{"is", "ISO", "8601"}},
		{"Win2000", []string{"Win", "2000"} },
		{"Bose QC35", []string{"Bose", "QC", "35"}},
		{"YAMLParser", []string{"YAML", "Parser"}},
		{"SOME_CONSTANT_STRING_REPRESENTATION", []string{"SOME", "CONSTANT", "STRING", "REPRESENTATION"}},
		{"joe, johnathan & john", []string{"joe", "johnathan", "john"}},
		{"a small-town family-owned business", []string{"a", "small", "town", "family", "owned", "business"}},
		{"-any-day-now-", []string{"any", "day", "now"}},
		{"a lot    of spaces   ", []string{"a", "lot", "of", "spaces"}},
		{"AnUnknownChar𖡄", []string{"An", "Unknown", "Char", "𖡄"}},
		{"�����������������������������", []string{}},
		{"invalidUTF8\xc5z", []string{"invalidUTF8\xc5z"}},
	}

	runExtractTest(t, tests)
}

func TestExtractWithOptionHyphenatedWords(t *testing.T) {
	tests := []testSet{
		{"-hyphenated-words", []string{"hyphenated-words"}},
		{"a later -hyphenated-word", []string{"a", "later", "hyphenated-word"}},
		{"a small-sized, dog-friendly, vacation home", []string{"a", "small-sized", "dog-friendly", "vacation", "home"}},
		{"other.chars_should-still*be>processed", []string{"other", "chars", "should-still", "be", "processed"}},
		{"-.-", []string{}},
		{"----------------", []string{}},
		{"----------a-b------------", []string{"a-b"}},
		{"-z-----------b", []string{"z", "b"}},
		{"a family-SIZED meal", []string{"a", "family", "SIZED", "meal"}},
	}

	runExtractTest(t, tests, words.AllowHyphenatedWords())
}

func TestExtractWithOptionIncludeSpace(t *testing.T) {
	tests := []testSet{
		{"a string with spaces", []string{"a", " ", "string", " ", "with", " ", "spaces"}},
		{"So   many   spaces", []string{"So", "   ", "many", "   ", "spaces"}},
		{"Spaces & Symbols", []string{"Spaces", " ", " ", "Symbols"}},
	}

	runExtractTest(t, tests, words.IncludeSpaces())
}

func TestExtractWithOptionIncludeSymbols(t *testing.T) {
	tests := []testSet{
		{"should>yield|any<symbol", []string{"should", ">", "yield", "|", "any", "<", "symbol"}},
		{"no punctuation!", []string{"no", "punctuation"}},
		{"<<<<<hi>>>>>", []string{"<<<<<", "hi", ">>>>>"}},
	}

	runExtractTest(t, tests, words.IncludeSymbols())
}

func TestExtractWithOptionIncludePunctuation(t *testing.T) {
	tests := []testSet{
		{"keep. all, punctuation!", []string{"keep", ".", "all", ",", "punctuation", "!"}},
		{">!..oops", []string{"!..", "oops"}},
	}

	runExtractTest(t, tests, words.IncludePunctuation())
}