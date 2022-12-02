package nlp

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

type TokCase struct {
	Text   string
	Tokens []string
}

func loadTokenizeCases(t *testing.T) []TokCase {
	file, err := os.Open("testdata/tokenize_cases.yml")
	require.NoError(t, err)
	defer file.Close()

	var tc []TokCase
	err = yaml.NewDecoder(file).Decode(&tc)
	require.NoError(t, err)
	return tc
}

// Exercise: Read test cases from tokenize_cases.yml
// instead of in-memory testCases
func TestTokenizeTable(t *testing.T) {
	/*
		testCases := []struct { // anonymous struct
			text   string
			tokens []string
		}{
			{"Who's on first?", []string{"who", "s", "on", "first"}},
			{"What's on second?", []string{"what", "s", "on", "second"}},
			{"", nil},
		}
	*/

	for _, tc := range loadTokenizeCases(t) {
		name := tc.Text
		if name == "" {
			name = "<empty>"
		}
		t.Run(name, func(t *testing.T) {
			tokens := Tokenize(tc.Text)
			require.Equal(t, tc.Tokens, tokens)
		})
	}

}

func TestTokenize(t *testing.T) {
	text := "Who's on first?"
	tokens := Tokenize(text)
	expected := []string{"who", "on", "first"}

	require.Equal(t, expected, tokens)

	/* before testify
	if !reflect.DeepEqual(expected, tokens) {
		t.Fatalf("expected: #%v, got: %#v", expected, tokens)
	}
	*/
}

func TestInCI(t *testing.T) {
	// Use "BUILD_NUMBER" in Jenkins
	if os.Getenv("CI") == "" {
		t.Skip("not in CI")
	}
}
