package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListPokemons(location string) (LocationArea, error) {
	url := baseURL + "/location-area/" + location

	if val, ok := c.cache.Get(url); ok {
		locationResp := LocationArea{}
		err := json.Unmarshal(val, &locationResp)
		if err != nil {
			return LocationArea{}, err
		}

		return locationResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationArea{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return LocationArea{}, fmt.Errorf("API returned status %d", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationArea{}, err
	}

	locationResp := LocationArea{}
	err = json.Unmarshal(data, &locationResp)
	if err != nil {
		return LocationArea{}, err
	}

	c.cache.Add(url, data)
	return locationResp, nil
}
