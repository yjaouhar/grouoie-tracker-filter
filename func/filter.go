package groupie

import (
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

var (
	min      int
	max      int
	start    int
	end      int
	first    int
	member   []string
	location string
	err      error
	lent     int
)

func Filter(w http.ResponseWriter, r *http.Request) {
	min, err = strconv.Atoi(r.FormValue("Min"))
	max, err = strconv.Atoi(r.FormValue("Max"))
	start, err = strconv.Atoi(r.FormValue("Start"))
	end, err = strconv.Atoi(r.FormValue("End"))
	member = r.Form["Members"]
	location = r.FormValue("Location")
	location = strings.ReplaceAll(location, ", ", "-")
	Result.Found = false
	if err != nil || min > max || start > end {
		Result.Found = true
	}
	Result.Searched = nil
	lent = len(Result.Searched)
	for i, v := range Result.Artist {
		if v.CreationDate >= min && v.CreationDate <= max {
			Result.Searched = append(Result.Searched, Result.Artist[i])
		}
	}

	lent = len(Result.Searched)
	for i, v := range Result.Searched {
		first, err = strconv.Atoi(v.FirstAlbum[len(v.FirstAlbum)-4:])
		if err != nil {
			Result.Found = true
		}
		if first >= start && first <= end {
			Result.Searched = append(Result.Searched, Result.Artist[i])
		}
	}
	Result.Searched = Result.Searched[lent:]

	if len(member) > 0 {
		lent = len(Result.Searched)
		for _, v := range Result.Searched {
			for _, x := range member {
				val, err := strconv.Atoi(x)
				if err != nil {
					Result.Found = true
				}
				if len(v.Members) == val {
					Result.Searched = append(Result.Searched, v)
				}
			}
		}
		Result.Searched = Result.Searched[lent:]
	}

	if len(location) > 0 {
		lent = len(Result.Searched)
		// for _, v := range Result.Searched {
		// 	for _, l := range v.Loco {
		// 		if strings.Contains((strings.ToLower(l)), strings.ToLower(location)) {
		// 			Result.Searched = append(Result.Searched, v)
		// 			break
		// 		}
		// 	}

		// }
		Result.Searched = Result.Searched[lent:]
	}
	if len(Result.Searched) == 0 {
		Result.Found = true
	}
	temp, err := template.ParseFiles("template/index.html")
	if err != nil {
		Error(w, http.StatusInternalServerError)
		return
	}
	// Execute the parsed template and write it to the response writer.
	ExecuteTemplate(temp, "display", w, nil, 0)
}
