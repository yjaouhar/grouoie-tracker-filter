package main

import (
	"fmt"
	"log"
	"net/http"

	groupie "groupie-tracker/func"
)

func main() {
	groupie.Start()
	http.HandleFunc("/search", groupie.Search)
	http.HandleFunc("/filter", groupie.Filter)
	http.HandleFunc("/style/{file}", groupie.Style)
	http.HandleFunc("/", groupie.Home)
	http.HandleFunc("/artist/{id}", groupie.ArtistInfo)
	fmt.Println("http://localhost:8060")
	log.Fatal(http.ListenAndServe(":8060", nil))
}
