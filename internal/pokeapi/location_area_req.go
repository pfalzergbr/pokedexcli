package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageUrl *string) (LocationAreasResp, error) {
	// Put together a new url


	endpoint := "/location-area"
	url := baseURL + endpoint

	if pageUrl != nil {
		url = *pageUrl
	}
	// Create the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreasResp{}, err
	}

	// Execute the request by the client passed in
	resp, err := c.httpClient.Do(req)

	if err != nil {
		return LocationAreasResp{}, err
	}
	// Close the response body once alld done
	defer resp.Body.Close()

	// 400+ status code is an error, so we throw a custom one here
	if resp.StatusCode > 399 {
		return LocationAreasResp{}, fmt.Errorf("bad status code %v", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResp{}, err
	}

	locationAreasResp := LocationAreasResp{}
	// Unmarshal the data into a LocationAreasResp struct
	err = json.Unmarshal(data, &locationAreasResp)
	if err != nil {
		return LocationAreasResp{}, err
	}

	return locationAreasResp, nil
}
