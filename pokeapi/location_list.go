package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageUrl *string) (RespLocationsArea, error) {
	url := baseUrl + "/locations-area"

	if pageUrl != nil {
		url = *pageUrl
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocationsArea{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocationsArea{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespLocationsArea{}, err
	}

	marshalled := RespLocationsArea{}
	if err := json.Unmarshal(data, &marshalled); err != nil {
		return RespLocationsArea{}, err
	}

	return marshalled, nil
}
