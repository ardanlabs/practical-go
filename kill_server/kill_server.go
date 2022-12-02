package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func main() {
	if err := KillServer("app.pid"); err != nil {
		// See also errors.Is & errors.As
		for e := err; e != nil; e = errors.Unwrap(e) {
			fmt.Println(e)
		}
		log.Fatalf("error: %s", err)
	}
}

func KillServer(pidFile string) error {
	// Go idiom: acquire resource, check for error, defer release
	file, err := os.Open(pidFile) // see also os.Create & os.OpenFile
	if err != nil {
		return err
	}
	defer file.Close()

	var pid int
	if _, err := fmt.Fscanf(file, "%d", &pid); err != nil {
		// %w will "wrap" the original error
		return fmt.Errorf("%s: bad pid - %w", pidFile, err)
	}

	fmt.Printf("killing %d\n", pid) // simulate kill

	// Design decision: What to do on error here?
	if err := os.Remove(pidFile); err != nil {
		log.Printf("warning: can't delete %s - %s", pidFile, err)
	}
	return nil
}
