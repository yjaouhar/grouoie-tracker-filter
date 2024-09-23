package groupie

import (
	"html/template"
	"net/http"
)

// Error handles rendering an error page with the specified HTTP status code
func Error(w http.ResponseWriter, status int) {
	tmp, err := template.ParseFiles("template/error.html")
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	// Create an instance of Err with the provided status code and corresponding message
	st_temp := Err{
		Status:  status,
		Message: http.StatusText(status),
	}
	ExecuteTemplate(tmp, "err", w, st_temp, status)
}
