package main

import (
	"context"
	"fmt"
	"strings"
	"time"
)

func main() {
	// We have 50 msec to return an answer
	ctx, cancel :=  // TODO
	defer cancel()
	url := "https://go.dev" // return the 7¢ ad
	// url := "http://go.dev" // return the default ad
	bid := bidOn(ctx, url)
	fmt.Println(bid)
}

// If algo didn't finish in time, return a default bid
func bidOn(ctx context.Context, url string) Bid {
	// TODO
	return Bid{}
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
