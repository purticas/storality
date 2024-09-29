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

func pageView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ViewPage"))
}

func pageCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create Page"))
}

func main() {
	fmt.Printf("Runing server on %s\n", port)
	mux := http.NewServeMux()
	mux.HandleFunc("/{$}", index)
	mux.HandleFunc("/page/view", pageView)
	mux.HandleFunc("/page/create", pageCreate)

	err := http.ListenAndServe(port, mux)
	log.Fatal(err)
}