package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	urls := []string{
		"https://www.google.com",
		"https://www.apple.com",
		"https://httpbin.org/status/404",
	}

	if err := siteTimes(urls); err != nil {
		fmt.Println("ERROR:", err)
	}
}

func siteTimes(urls []string) error {
	// TODO
	return nil
}

func siteTime(url string) (time.Duration, error) {
	start := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("bad status: %d %s", resp.StatusCode, resp.Status)
	}
	_, err = io.Copy(io.Discard, resp.Body)
	if err != nil {
		return 0, err
	}

	return time.Since(start), nil
}
