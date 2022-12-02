package main

import (
	"encoding/json"
	"expvar"
	"os"

	// _ "expvar"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/ardanlabs/practical-go/nlp"
	"github.com/ardanlabs/practical-go/nlp/stemmer"
)

// Exercise: Write a tokenizeHandler that will get the text in the request body
// curl -d "Who's on first?" http://localhost:8080/tokenize

var (
	tokenizeCalls = expvar.NewInt("tokenize.calls")
)

// config: defaults < config file < environment < command line
// default: code
// config file: yaml.v2, toml ...
// environ: os.Getenv
// command line: flag
// Validate configuration before applications starts

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/health", healthHandler).Methods("GET")
	r.HandleFunc("/tokenize", tokenizeHandler).Methods("POST")
	r.HandleFunc("/stem/{word}", stemHandler).Methods("GET")
	http.Handle("/", r)
	// built-in routing: if path ends with / - prefix math, otherwise - exact match

	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = ":8080"
	}

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("INFO: server starting on %s", addr)

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("error: %s", err)
	}
}

func stemHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	word := vars["word"]
	stem := stemmer.Stem(word)
	fmt.Fprintln(w, stem)
}

func tokenizeHandler(w http.ResponseWriter, r *http.Request) {
	// Step 1: Get, unmarshal & validate request
	/* Before gorilla/mux
	if r.Method != http.MethodPost {
		http.Error(w, "bad method", http.StatusMethodNotAllowed)
		return
	}
	*/

	tokenizeCalls.Add(1)

	const MB = 1 << 20
	rdr := io.LimitReader(r.Body, 3*MB)
	data, err := io.ReadAll(rdr)
	if err != nil {
		http.Error(w, "can't read", http.StatusBadRequest)
		return
	}

	// Step 2: Work
	tokens := nlp.Tokenize(string(data))

	// Step 3: Marshal response
	resp := map[string]any{
		"tokens": tokens,
	}

	body, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, "can't marshal to JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Check database connection ...
	fmt.Fprintln(w, "OK")
}
