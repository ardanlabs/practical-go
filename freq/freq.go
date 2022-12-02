package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

var (
	// `` is a "raw" string: \ is just a \ and you can have multiline strings

	// "Who's on first?" -> [Who s on first]
	wordRe = regexp.MustCompile(`[a-zA-Z]+`)
)

func main() {
	file, err := os.Open("sherlock.txt")
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	defer file.Close()

	// mapDemo()
	freq, err := frequency(file)
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	fmt.Println(maxWord(freq))
}

// What is the most common word, ignoring case, in sherlock.txt?
// Word frequency

func maxWord(freq map[string]int) string {
	maxW, maxN := "", 0
	for w, n := range freq {
		if n > maxN {
			maxW, maxN = w, n
		}
	}
	return maxW
}

// frequency return map of word -> count (ignoring case)
func frequency(r io.Reader) (map[string]int, error) {
	s := bufio.NewScanner(r)
	freq := make(map[string]int)
	for s.Scan() {
		words := wordRe.FindAllString(s.Text(), -1)
		for _, w := range words {
			freq[strings.ToLower(w)]++
		}
	}
	if err := s.Err(); err != nil {
		return nil, err
	}
	return freq, nil
}

func mapDemo() {
	// var stocks map[string]float64 // panic on add

	stocks := map[string]float64{
		"AAPL": 1.2,
		"IBM":  3.4,
	}
	price := stocks["GOOG"]
	fmt.Println(price) // 0
	price, ok := stocks["GOOG"]
	if ok {
		fmt.Println(price)
	} else {
		fmt.Println("GOOG not found")
	}

	if _, ok := stocks["GOOG"]; !ok {
		fmt.Println("GOOG not found")
	}

	stocks["GOOG"] = 5.6 // add

	for s := range stocks { // keys
		fmt.Println(s)
	}

	for s, p := range stocks { // key + value
		fmt.Println(s, "->", p)
	}
}
