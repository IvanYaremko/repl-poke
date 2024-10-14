package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) PokemonDetails(name string) (Pokemon, error) {
	url := baseUrl + "/pokemon/" + name

	cacheData, ok := c.cache.Get(url)
	if ok {
		pokemon := Pokemon{}
		if err := json.Unmarshal(cacheData, &pokemon); err != nil {
			return Pokemon{}, fmt.Errorf("pokemon details cache unmarshalling %w: ", err)
		}

		return pokemon, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, fmt.Errorf("pokemon details creating request %w: ", err)
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, fmt.Errorf("getting pokmon request %w: ", err)
	}

	byteData, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, fmt.Errorf("pokemone details io read all %w: ", err)
	}

	pokemon := Pokemon{}
	if err := json.Unmarshal(byteData, &pokemon); err != nil {
		return Pokemon{}, fmt.Errorf("pokemon details unmarshalling %w: ", err)
	}

	c.cache.Add(url, byteData)
	return pokemon, nil
}
