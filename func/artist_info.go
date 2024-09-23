package groupie

import (
	"html/template"
	"net/http"
	"strconv"
)

// Posthandel handles POST requests to display artist information based on the provided artist ID
func ArtistInfo(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	IDa, er := strconv.Atoi(id)
	if er != nil || IDa <= 0 || IDa > 52 {
		Error(w, http.StatusNotFound)
		return
	}

	if r.URL.Path != "/artist/"+id {
		Error(w, http.StatusNotFound)
		return
	}
	temp, err := template.ParseFiles("template/index2.html")
	if err != nil {
		Error(w, http.StatusInternalServerError)
		return
	}
	// Clear existing relations in the Cards.Rela.Relation map
	for key := range Cards.Rela.Relation {
		delete(Cards.Rela.Relation, key)
	}
	Cards.Art = Result.Tbn[IDa-1]
	if !Fetch("locations", "/"+id) {
		Error(w, http.StatusInternalServerError)
		return
	}
	if !Fetch("dates", "/"+id) {
		Error(w, http.StatusInternalServerError)
		return
	}
	if !Fetch("relation", "/"+id) {
		Error(w, http.StatusInternalServerError)
		return
	}

	// Execute the parsed template and write it to the response writer
	ExecuteTemplate(temp, "", w, nil, 0)
}
