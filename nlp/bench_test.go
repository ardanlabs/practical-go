package nlp

import "testing"

const (
	text = "Design the architecture, name the components, document the details."
)

func BenchmarkTokenize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tokens := Tokenize(text)
		if len(tokens) != 9 {
			b.Fatal(tokens)
		}
	}
}
