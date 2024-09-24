package groupie

var (
	Isfetched bool
	Result    Rounder
	Uniq        = make(map[string]interface{})
)

const Url = "https://groupietrackers.herokuapp.com/api/"

type Rounder struct {
	Artist []struct {
		ID           int      `json:"id"`
		Image        string   `json:"image"`
		Name         string   `json:"name"`
		Members      []string `json:"members"`
		CreationDate int      `json:"creationDate"`
		FirstAlbum   string   `json:"firstAlbum"`
	}
	Searched []struct {
		ID           int      `json:"id"`
		Image        string   `json:"image"`
		Name         string   `json:"name"`
		Members      []string `json:"members"`
		CreationDate int      `json:"creationDate"`
		FirstAlbum   string   `json:"firstAlbum"`
	}

	Relation struct { // struct that groups 
		Index []struct {
			AllRelations map[string][]string `json:"datesLocations"`
		} `json:"index"`
	}
	ArtRelation map[string][]string

	UniMap map[string]interface{}// map[string]map[string]bool
	Found bool
}
type Err struct {
	Status  int
	Message string
}
