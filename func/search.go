package groupie

import (
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

// Search handles search requests for artists based on query parameters
func Search(w http.ResponseWriter, r *http.Request) {
	Result.Found = false  // Initialize the search result as not found
	Result.Searched = nil // Clear any previous search results
	query := r.URL.Query().Get("art")
	if query == "" {
		Error(w, 404)
		return
	}

	// Iterate over the list of artists to find matches
	for i, v := range Result.Artist {
		if strings.Contains(strings.ToLower(v.Name), strings.ToLower(query)) {
			Result.Searched = append(Result.Searched, Result.Artist[i])
			continue
		} else if strings.Contains(strings.ToLower(v.FirstAlbum), strings.ToLower(query)) && len(query) != 4 {
			Result.Searched = append(Result.Searched, Result.Artist[i])
			continue
		} else if strings.Contains(strings.ToLower(strconv.Itoa(v.CreationDate)), strings.ToLower(query)) && len(query) == 4 {
			Result.Searched = append(Result.Searched, Result.Artist[i])
			continue
		} else {
			found := false
			for _, mem := range v.Members {
				if strings.Contains(strings.ToLower(mem), strings.ToLower(query)) {
					Result.Searched = append(Result.Searched, Result.Artist[i])
					found = true
					break
				}
			}
			// If no member match is found, check relations
			if !found {
				for v := range Result.Relation.Index[i].AllRelations {
					if strings.Contains(strings.ToLower(v), strings.ToLower(query)) {
						Result.Searched = append(Result.Searched, Result.Artist[i])
						break
					}
				}
			}

		}
	}
	// If no results were found, set the found flag to true
	if len(Result.Searched) == 0 {
		Result.Found = true
	}

	temp, err := template.ParseFiles("template/index.html")
	if err != nil {
		Error(w, http.StatusInternalServerError)
		return
	}

	ExecuteTemplate(temp, "display", w, nil, 0)
}
