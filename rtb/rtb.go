package main

import (
	"context"
	"fmt"
	"strings"
	"time"
)

func main() {
	// We have 50 msec to return an answer
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()
	url := "https://go.dev"
	bid := bidOn(ctx, url)
	fmt.Println(bid)
}

// If algo didn't finish in time, return a default bid
func bidOn(ctx context.Context, url string) Bid {
	ch := make(chan Bid, 1) // buffered channel to avoid goroutine leak
	go func() {
		ch <- bestBid(url)
	}()

	select {
	case bid := <-ch:
		return bid
	case <-ctx.Done():
		return defaultBid
	}
}

var defaultBid = Bid{
	AdURL: "http://adsЯus.com/default",
	Price: 3,
}

// Written by Algo team, time to completion varies
func bestBid(url string) Bid {
	// Simulate work
	d := 100 * time.Millisecond
	if strings.HasPrefix(url, "https://") {
		d = 20 * time.Millisecond
	}
	time.Sleep(d)

	return Bid{
		AdURL: "http://adsЯus.com/ad17",
		Price: 7,
	}
}

type Bid struct {
	AdURL string
	Price int // In ¢
}
