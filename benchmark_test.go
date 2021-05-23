package words_test

import (
	"github.com/imbue11235/words"
	"testing"
)

func benchmarkExtract(input string, b *testing.B) {
	for i := 0; i < b.N; i++ {
		words.Extract(input)
	}
}

func BenchmarkExtract1RuneType(b *testing.B) {
	benchmarkExtract("a sentence", b)
}

func BenchmarkExtract2RuneTypes(b *testing.B) {
	benchmarkExtract("Windows are Windows", b)
}

func BenchmarkExtract3RuneTypes(b *testing.B) {
	benchmarkExtract("Apple Iphone 12 Pro Max", b)
}