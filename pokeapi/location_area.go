package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) LocationArea(locationId string) (ResLocationsArea, error) {
	cachedData, ok := c.cache.Get(locationId)
	if ok {
		areaData := ResLocationsArea{}
		if err := json.Unmarshal(cachedData, &areaData); err != nil {
			return ResLocationsArea{}, fmt.Errorf("unmarshall cache data error: %w", err)
		}
		return areaData, nil
	}

	url := baseUrl + "/location-area/" + locationId

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ResLocationsArea{}, fmt.Errorf("new request error: %w", err)
	}

	response, err := c.httpClient.Do(req)
	if err != nil {
		return ResLocationsArea{}, fmt.Errorf("get data error: %w", err)
	}

	byteData, err := io.ReadAll(response.Body)
	if err != nil {
		return ResLocationsArea{}, fmt.Errorf("io read data error: %w", err)
	}

	areaData := ResLocationsArea{}
	if err := json.Unmarshal(byteData, &areaData); err != nil {
		return ResLocationsArea{}, fmt.Errorf("unmarshall data error: %w", err)
	}

	c.cache.Add(locationId, byteData)
	return areaData, nil
}
