package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func handleMyGreet(w http.ResponseWriter, r *http.Request) {
	greet(w, "World")
}

func main() {
	http.HandleFunc("/", handleMyGreet)
	log.Fatal(http.ListenAndServe(":5000", nil))
}
