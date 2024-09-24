package groupie

// Fetch artist data and store the result in Isfetched
func Start() {
	Isfetched = Fetch("artists")
	Isfetched = Fetch("relation")
	Uni()
	Result.UniMap = Uniq
}
