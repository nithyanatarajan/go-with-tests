package main

import (
	"log"
	"net/http"
	"os"

	"github.com/nithyanatarajan/go-with-tests/print"
)

func main() {
	print.Greet(os.Stdout, "World!")
	webGreeter := func(w http.ResponseWriter, request *http.Request) {
		print.Greet(w, "World!")
	}
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(webGreeter)))
}
