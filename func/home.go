package groupie

import (
	"html/template"
	"net/http"
)

// Gethandel handles HTTP requests for the home page ("/")
func Home(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		Error(w, http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != "/" {
		Error(w, http.StatusNotFound)
		return
	}

	// Check if data has been fetched and is available for rendering
	if !Isfetched {
		Error(w, http.StatusInternalServerError)
		return
	}
	Result.Found = false
	temp, err := template.ParseFiles("template/index.html")
	if err != nil {
		Error(w, http.StatusInternalServerError)
		return
	}

	Result.Searched = Result.Artist

	ExecuteTemplate(temp, "display", w, nil, 0)
}
