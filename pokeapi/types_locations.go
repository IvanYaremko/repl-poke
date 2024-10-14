package pokeapi

type RespLocationsArea struct {
	Count    int                        `json:"count"`
	Next     string                     `json:"next"`
	Previous string                     `json:"previous"`
	Results  []RespLocationsAreaResults `json:"results"`
}

type RespLocationsAreaResults struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
