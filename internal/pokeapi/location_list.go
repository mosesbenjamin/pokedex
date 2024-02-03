package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// ListLocations-
func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}
	if val, ok := c.cache.Get(url); ok {
		locationResponse := RespShallowLocations{}
		err := json.Unmarshal(val, &locationResponse)
		if err != nil {
			return RespShallowLocations{}, err
		}
		return locationResponse, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}

	locationResponse := RespShallowLocations{}
	err = json.Unmarshal(data, &locationResponse)
	if err != nil {
		return RespShallowLocations{}, err
	}
	c.cache.Add(url, data)
	return locationResponse, nil
}
