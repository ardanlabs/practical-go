package main

import (
	"encoding/json"
	"expvar"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/353solutions/nlp"
	"github.com/353solutions/nlp/stemmer"
)

var (
	numTok = expvar.NewInt("tokenize.calls")
)

func main() {

	// Create server
	logger := log.New(log.Writer(), "[nlpd] ", log.LstdFlags|log.Lshortfile)
	s := Server{
		logger: logger, //  dependency injection
	}
	// routing
	// /health is an exact match
	// /health/ is a prefix match
	/*
		http.HandleFunc("/health", healthHandler)
		http.HandleFunc("/tokenize", tokenizeHandler)
	*/
	r := mux.NewRouter()
	r.HandleFunc("/health", s.healthHandler).Methods(http.MethodGet)
	r.HandleFunc("/tokenize", s.tokenizeHandler).Methods(http.MethodPost)
	r.HandleFunc("/stem/{word}", s.stemHandler).Methods(http.MethodGet)
	http.Handle("/", r)

	// run server
	addr := os.Getenv("NLPD_ADDR")
	if addr == "" {
		addr = ":8080"
	}
	s.logger.Printf("server starting on %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("error: %s", err)
	}
}

type Server struct {
	logger *log.Logger
}

func (s *Server) stemHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	word := vars["word"]
	stem := stemmer.Stem(word)
	fmt.Fprintln(w, stem)
}

func (s *Server) tokenizeHandler(w http.ResponseWriter, r *http.Request) {
	/* Before gorilla/mux
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	*/

	numTok.Add(1)

	// Step 1: Get, convert & validate the data
	rdr := io.LimitReader(r.Body, 1_000_000)
	data, err := io.ReadAll(rdr)
	if err != nil {
		s.logger.Printf("error: can't read - %s", err)
		http.Error(w, "can't read", http.StatusBadRequest)
		return
	}

	if len(data) == 0 {
		http.Error(w, "missing data", http.StatusBadRequest)
		return
	}

	text := string(data)

	// Step 2: Work
	tokens := nlp.Tokenize(text)

	// Step 3: Encode & emit output
	resp := map[string]any{
		"tokens": tokens,
	}
	// You can also do:
	// err = json.NewEncoder(w).Encode(resp)
	data, err = json.Marshal(resp)
	if err != nil {
		http.Error(w, "can't encode", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

// exercise: Write tokenizeHandler that will read the text form r.Body and
// return JSON in the format: "{"tokens": ["who", "on", "first"]}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Run a health check
	fmt.Fprintln(w, "OK")
}
