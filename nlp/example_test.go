package nlp_test

import (
	"fmt"

	"github.com/353solutions/nlp"
)

func ExampleTokenize() {
	text := "Who's on first?"
	tokens := nlp.Tokenize(text)
	fmt.Println(tokens)

	// Output:
	// [who on first]
}

/*
Test discovery:
For every file ending with _test.go, run every function that matches either:
- Example[A-Z_].*, body must include // Output: comment
- Test[A-Z_].*
*/
