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

func makeExtractTest(t *testing.T, tests []testSet, options ...words.Option) {
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
	}

	makeExtractTest(t, tests)
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
	}

	makeExtractTest(t, tests, words.AllowHyphenatedWords())
}