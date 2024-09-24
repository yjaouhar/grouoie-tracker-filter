package groupie

import (
	"bytes"
	"html/template"
	"net/http"
)

// ExecuteTemplate renders a template with the provided data and writes it to the response writer.
func ExecuteTemplate(temp *template.Template, s string, w http.ResponseWriter, info interface{}, status int) {
	var buf bytes.Buffer // Buffer to hold the rendered template content temporarily
	var err error

	// Handle different template display cases based on the value of 's'
	if s == "display" {
		err = temp.Execute(&buf, Result)
		if err != nil {
			Error(w, http.StatusInternalServerError)
			return
		}
		temp.Execute(w, Result)

	} else if s == "err" {
		err = temp.Execute(&buf, info)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(status)
		temp.Execute(w, info)
	} else {
		ArtInfo := Rounder{ //[1]
			Artist:      Result.Artist[status-1 : status],
			ArtRelation: Result.Relation.Index[status-1].AllRelations,
		}

		err = temp.Execute(&buf, ArtInfo)
		if err != nil {
			Error(w, http.StatusInternalServerError)
			return
		}
		temp.Execute(w, ArtInfo)
	}
}
