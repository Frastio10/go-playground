package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/a-h/templ"
)

func main() {
	component := home("Frastio")
	http.Handle("/", templ.Handler(component))

	fmt.Println("Listening on :3000")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal("Failed to listen: ", err)

	}
}
