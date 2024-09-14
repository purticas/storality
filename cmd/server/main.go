package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = ":3000"

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Storality CMS"))
}

func main() {
	fmt.Printf("Runing server on %s\n", port)
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)

	err := http.ListenAndServe(port, mux)
	log.Fatal(err)
}