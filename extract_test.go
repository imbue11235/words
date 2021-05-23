package words_test

import (
	"github.com/imbue11235/words"
	"reflect"
	"testing"
)

func TestExtract(t *testing.T) {
	tests := []struct{
		input string
		expected []string
	} {
		{"Win2000", []string{"Win", "2000"} },
		{"YAMLParser", []string{"YAML", "Parser"}},
		{"SOME_CONSTANT_STRING_REPRESENTATION", []string{"SOME", "CONSTANT", "STRING", "REPRESENTATION"}},
	}

	for _, test := range tests {
		extraction := words.Extract(test.input)
		if !reflect.DeepEqual(extraction, test.expected) {
			t.Errorf("Expected %v to be %v", extraction, test.expected)
		}
	}
}
