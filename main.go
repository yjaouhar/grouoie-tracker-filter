package main

import (
	"fmt"
	"log"
	"net/http"

	groupie "groupie-tracker/func"
)

func main() {
	//var mo interface{}
	groupie.Isfetched = groupie.Fetch("artists", "")
	groupie.Isfetched = groupie.Fetch("location", "s")
	http.HandleFunc("/search", groupie.Search)
	http.HandleFunc("/filter", groupie.Filter)
	http.HandleFunc("/style/{file}", groupie.Style)
	http.HandleFunc("/", groupie.Home)
	http.HandleFunc("/artist/{id}", groupie.ArtistInfo)
	fmt.Println("http://localhost:8089")
	log.Fatal(http.ListenAndServe(":8089", nil))
}
