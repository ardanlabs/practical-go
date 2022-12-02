package nlp

import (
	"regexp"
	"strings"

	"github.com/ardanlabs/practical-go/nlp/stemmer"
)

var (
	wordRe = regexp.MustCompile(`[a-zA-Z]+`)
)

// Tokenize returns slice of tokens (lower case) found in text.
func Tokenize(text string) []string {
	words := wordRe.FindAllString(text, -1)
	// var tokens []string
	tokens := make([]string, 0, 20) // 75 percentile of text has less than 20 tokens
	for _, w := range words {
		token := strings.ToLower(w)
		token = stemmer.Stem(token)
		if token != "" {
			tokens = append(tokens, token)
		}
	}
	return tokens
}
