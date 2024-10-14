package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageUrl *string) (RespLocationsArea, error) {
	url := baseUrl + "/location-area"

	if pageUrl != nil {
		url = *pageUrl
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocationsArea{}, fmt.Errorf("error creating request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocationsArea{}, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespLocationsArea{}, fmt.Errorf("error reading body: %w", err)
	}

	marshalled := RespLocationsArea{}
	err = json.Unmarshal(data, &marshalled)
	if err != nil {
		return RespLocationsArea{}, fmt.Errorf("error unmarshalling data: %w", err)
	}

	return marshalled, nil
}
