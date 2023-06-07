package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func httpResGreeter(w http.ResponseWriter, r *http.Request) {
	Greet(w, "CHUNGUS")
}

func main() {
	Greet(os.Stdout, "Quandale dingle")
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(httpResGreeter)))
}
