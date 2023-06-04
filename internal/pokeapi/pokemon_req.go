package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	// Put together a new url

	endpoint := "/pokemon/" + pokemonName
	url := baseURL + endpoint

	// On cache hit, we can get all the data straight away.
	if cachedVal, ok := c.cache.Get(url); ok {
		pokemon := Pokemon{}

		err := json.Unmarshal(cachedVal, &pokemon)

		if err != nil {
			return Pokemon{}, err
		}

		return pokemon, nil
	}

	// Create the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	// Execute the request by the client passed in
	resp, err := c.httpClient.Do(req)

	if err != nil {
		return Pokemon{}, err
	}
	// Close the response body once alld done
	defer resp.Body.Close()

	// 400+ status code is an error, so we throw a custom one here
	if resp.StatusCode > 399 {
		return Pokemon{}, fmt.Errorf("bad status code %v", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	pokemon := Pokemon{}

	// Add the successfully retrived data to the cache
	c.cache.Add(url, data)
	// Unmarshal the data into a PokemonsResp struct
	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}

	return pokemon, nil
}
