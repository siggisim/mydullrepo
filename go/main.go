package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", getRoot)

	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		log.Fatalf("failed to ListenAndServe: %s", err)
	}
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	log.Printf("got request")
	io.WriteString(w, "Hello, World!\n")
}
