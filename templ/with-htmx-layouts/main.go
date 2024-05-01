package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
)

func main() {
	http.Handle("/", templ.Handler(Index()))
	http.Handle("/about", templ.Handler(About()))

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}
