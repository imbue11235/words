package words_test

import (
	"fmt"
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
		{"joe, johnathan & john", []string{"joe", "johnathan", "john"}},
		{"a small-town family-owned business", []string{"a", "small", "town", "family", "owned", "business"}},
	}

	for _, test := range tests {
		extraction := words.Extract(test.input)
		fmt.Println(test.input, len(extraction), len(test.expected))
		if !reflect.DeepEqual(extraction, test.expected) {
			t.Errorf("Expected %v to be %v", extraction, test.expected)
		}
	}
}