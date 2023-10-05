package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Starting Serer...")

	http.HandleFunc("/", getPage)
	http.HandleFunc("/add-film/", postFilm)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
