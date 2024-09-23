package groupie

import (
	"net/http"
	"os"
)

// Style handles requests for style files, serving them if they exist
func Style(w http.ResponseWriter, r *http.Request) {
	// Extract the file name from the request URL path
	f := r.PathValue("file")

	style := http.StripPrefix("/style/", http.FileServer(http.Dir("style"))) // Create a file server that serves files from the "style" directory
	// Check if the requested file exists by trying to read it
	_, err := os.ReadFile("style/" + f)

	if err != nil {
		Error(w, http.StatusNotFound)
		return
	}

	style.ServeHTTP(w, r)
}
