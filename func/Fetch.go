package groupie

import (
	"encoding/json"
	"io"
	"net/http"
)

// Fetch retrieves data from a specified URL  and unmarshals it into appropriate variables based on the type of data requested
func Fetch(s, id string) bool {
	response, err := http.Get(Url + s + id)
	if err != nil {
		return false
	}
	defer response.Body.Close() // Ensure the response body is closed when the function exits.

	if response.StatusCode != http.StatusOK {
		return false
	}
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return false
	}
	// Unmarshal the JSON data into the appropriate variable based on the type of data requested.
	if s == "artists" {
		err = json.Unmarshal(data, &Result.Tbn)

	} else if s == "locations" {
		err = json.Unmarshal(data, &Cards.Loca)
	} else if s == "dates" {
		err = json.Unmarshal(data, &Cards.Conc)
	} else if s == "relation" {
		err = json.Unmarshal(data, &Cards.Rela)
	} else {
		err = json.Unmarshal(data, &Mok)
	}
	if err != nil {
		return false
	}
	return true
}
