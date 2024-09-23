package groupie

var (
	Artist      Artists
	Location    Locations
	ConcertDate ConcertDates
	Relation    Relations
	Cards       Card
	Fetched     bool
	Id          int
	Isfetched   bool
	Mok         blasa
	Result      jdak
)

type Err struct {
	Status  int
	Message string
}

type Artists []struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Loco         []string
}
type mok []struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Loco         []string
}
type jdak struct {
	Tbn []struct {
		ID           int      `json:"id"`
		Image        string   `json:"image"`
		Name         string   `json:"name"`
		Members      []string `json:"members"`
		CreationDate int      `json:"creationDate"`
		FirstAlbum   string   `json:"firstAlbum"`
		Loco         []string
	}
	Mok []struct {
		ID           int      `json:"id"`
		Image        string   `json:"image"`
		Name         string   `json:"name"`
		Members      []string `json:"members"`
		CreationDate int      `json:"creationDate"`
		FirstAlbum   string   `json:"firstAlbum"`
		Loco         []string
	}
	F struct {
		Mp map[string]bool
	}
}

type blasa struct {
	Index []struct {
		//ID        int      `json:"id"`
		Locations []string `json:"locations"`
		//Dates     string   `json:"dates"`
	} `json:"index"`
}

type Locations struct {
	Locations []string `json:"locations"`
}
type ConcertDates struct {
	Dates []string `json:"dates"`
}

type Relations struct {
	Relation map[string][]string `json:"datesLocations"`
}
type Card struct {
	Art struct {
		ID           int      `json:"id"`
		Image        string   `json:"image"`
		Name         string   `json:"name"`
		Members      []string `json:"members"`
		CreationDate int      `json:"creationDate"`
		FirstAlbum   string   `json:"firstAlbum"`
		Loco         []string
	}
	Loca Locations
	Conc ConcertDates
	Rela Relations
}

const Url = "https://groupietrackers.herokuapp.com/api/"
