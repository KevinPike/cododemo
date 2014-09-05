package main

import (
	"log"
	"fmt"
	"net/http"
	"os"
	"html"
	)


func main() {
	//http.Handle("/foo", fooHandler)

	h, _ := os.Hostname()
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello, %q - %s", html.EscapeString(r.URL.Path), h)
			})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
