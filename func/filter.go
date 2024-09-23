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

	if err != nil || min > max || start > end {
		Error(w, 404)
		return
	}
	Result.Mok = nil
	lent = len(Result.Mok)
	for i, v := range Result.Tbn {
		if v.CreationDate >= min && v.CreationDate <= max {
			Result.Mok = append(Result.Mok, Result.Tbn[i])
		}
	}

	lent = len(Result.Mok)
	for i, v := range Result.Mok {
		first, err = strconv.Atoi(v.FirstAlbum[len(v.FirstAlbum)-4:])
		if err != nil {
			Error(w, 404)
			return
		}
		if first >= start && first <= end {
			Result.Mok = append(Result.Mok, Result.Mok[i])
		}
	}
	Result.Mok = Result.Mok[lent:]

	if len(member) > 0 {
		lent = len(Result.Mok)
		for _, v := range Result.Mok {
			for _, x := range member {
				val, err := strconv.Atoi(x)
				if err != nil {
					Error(w, 404)
					return
				}
				if len(v.Members) == val {
					Result.Mok = append(Result.Mok, v)
				}
			}
		}
		Result.Mok = Result.Mok[lent:]
	}

	if len(location) > 0 {
		lent = len(Result.Mok)
		for _, v := range Result.Mok {
			for _, l := range v.Loco {
				if strings.Contains((strings.ToLower(l)), strings.ToLower(location)) {
					Result.Mok = append(Result.Mok, v)
					break
				}
			}

		}
		Result.Mok = Result.Mok[lent:]
	}
	temp, err := template.ParseFiles("template/index.html")
	if err != nil {
		Error(w, http.StatusInternalServerError)
		return
	}
	// Execute the parsed template and write it to the response writer.
	ExecuteTemplate(temp, "alo", w, nil, 0)
}
