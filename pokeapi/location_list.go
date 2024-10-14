package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageUrl *string) (RespLocationsList, error) {
	url := baseUrl + "/location-area"

	if pageUrl != nil {
		url = *pageUrl
	}

	val, ok := c.cache.Get(url)
	if ok {
		response := RespLocationsList{}
		err := json.Unmarshal(val, &response)
		if err != nil {
			return RespLocationsList{}, err
		}

		return response, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocationsList{}, fmt.Errorf("error creating request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocationsList{}, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespLocationsList{}, fmt.Errorf("error reading body: %w", err)
	}

	response := RespLocationsList{}
	err = json.Unmarshal(data, &response)
	if err != nil {
		return RespLocationsList{}, fmt.Errorf("error unmarshalling data: %w", err)
	}

	return response, nil
}
