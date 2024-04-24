package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>hllo</h1>")
	fmt.Println("Endpoint")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Articles hit")
	json.NewEncoder(w).Encode(Articles)
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", returnAllArticles)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	Articles = []Article{
		Article{Title: "Gimme", Desc: "Descriptionr", Content: "Contentooo"},
		Article{Title: "Gimme", Desc: "Descriptionr", Content: "Contentooo"},
	}

	handleRequests()
}
