package main

import (
	"fmt"
	"unicode/utf8"
)

/*
banner("Go", 6)

  Go
------
*/

func main() {
	banner("Go", 6)
	banner("G☺", 6)

	fmt.Println("Go:", len("Go"))
	fmt.Println("G☺:", len("G☺")) // length in bytes

	s := "G☺"
	b := s[0]
	fmt.Printf("%c %T\n", b, b)
	fmt.Println(s[:1]) // slicing (half-open range), on bytes
	fmt.Println(s[1:])
	// fmt.Println(s[:2]) // will take G and first byte of ☺
	// s[0] = 'g' // string are immutable

	for i, c := range "G☺!" {
		fmt.Printf("%d: %c %T\n", i, c, c)
	}

	// byte = uint8
	// rune = int32 (code point, character)

	x, y := 1, "1"
	fmt.Printf("x=%v, y=%v\n", x, y)
	fmt.Printf("x=%#v, y=%#v\n", x, y)
	// fmt.Sprintf() // return new string
}

func banner(text string, width int) {
	// offset := (width - len(text)) / 2
	offset := (width - utf8.RuneCountInString(text)) / 2
	for i := 0; i < offset; i++ {
		fmt.Print(" ")
	}
	fmt.Println(text)

	for i := 0; i < width; i++ {
		fmt.Print("-")
	}
	fmt.Println()
}
