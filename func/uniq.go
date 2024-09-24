package groupie

import (
	"strconv"
)// map[string]map[string]bool

// Uni processes and populates unique data from the Result struct into the Uniq map.
func Uni() {
	temp := make(map[string]bool) 
	for _, v := range Result.Relation.Index { // Iterate over all relations to collect unique locations
		for loca := range v.AllRelations {
			if !temp[loca] {
				temp[loca] = true
			}
		}
	}
	Uniq["location"] = temp
	temp = make(map[string]bool)
	for _, v := range Result.Artist {	// Iterate over all artists to collect unique members
		for _, member := range v.Members {
			if !temp[member] {
				temp[member] = true
			}
		}
	}
	Uniq["member"] = temp
	temp = make(map[string]bool)
	for _, Arti := range Result.Artist {// Iterate over all artists to collect unique creation dates
		if !temp[strconv.Itoa(Arti.CreationDate)] {
			temp[strconv.Itoa(Arti.CreationDate)] = true
		}
	}
	Uniq["creation date"] = temp
	temp = make(map[string]bool)
	for _, Arti := range Result.Artist {// Iterate over all artists to collect unique first albums
		if !temp[Arti.FirstAlbum] {
			temp[Arti.FirstAlbum] = true
		}
	}
	Uniq["First Album"] = temp
}
