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
		{ "Windows2000", []string{"Windows", "2000"} },
	}

	for _, test := range tests {
		extraction := words.Extract(test.input)
		if !reflect.DeepEqual(extraction, test.expected) {
			t.Errorf("Expected %v to be %v", extraction, test.expected)
		}
	}
}
