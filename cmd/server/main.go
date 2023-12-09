package main

import (
	"net/http"

	"github.com/a-h/templ"
)

func main() {
	page := Index()

	http.Handle("/", templ.Handler(page))

	http.ListenAndServe(":3000", nil)
}
