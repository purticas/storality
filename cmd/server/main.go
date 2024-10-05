package main

import (
	"fmt"
	"log"
	"net/http"
	"storality/admin"
	"storality/app"
)

const port = ":3000"

func main() {
	router := http.NewServeMux()
	app := app.Init(router)
	admin.Serve(router)

	fmt.Printf("Runing server on %s\n", port)
	err := http.ListenAndServe(port, router)
	if(err != nil) {
		log.Fatal(err)
	}
}