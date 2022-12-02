package stemmer

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStemmer(t *testing.T) {
	cases := []struct {
		word string
		stem string
	}{
		{"works", "work"},
		{"working", "work"},
		{"worked", "work"},
		{"work", "work"},
	}

	for _, tc := range cases {
		t.Run(tc.word, func(t *testing.T) {
			stem := Stem(tc.word)
			require.Equal(t, tc.stem, stem)
		})
	}

}
