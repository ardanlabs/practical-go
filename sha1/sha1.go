/*
$ cat http.log.gz | gunzip | sha1sum
*/
package main

import (
	"compress/gzip"
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	// Goland
	// sig, err := fileSHA1("sha1/http.log.gz")
	sig, err := fileSHA1("http.log.gz")
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	fmt.Println(sig)

	sig, err = fileSHA1("sha1.go")
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	fmt.Println(sig)
}

func fileSHA1(fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var r io.Reader = file

	if strings.HasSuffix(fileName, ".gz") {
		var err error // shadow err from line 33
		r, err = gzip.NewReader(file)
		if err != nil {
			return "", err
		}
		// log.Printf("r: %v", r)
	}
	w := sha1.New()
	if _, err := io.Copy(w, r); err != nil {
		return "", err
	}
	data := w.Sum(nil)
	return fmt.Sprintf("%x", data), nil
}
