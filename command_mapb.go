package main

import (
	"fmt"
)

func commandMapb(cfg *config) error {
	if cfg.prevLocationAreasURL == nil {
		fmt.Println("No previous page!")
		return nil
	}
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.prevLocationAreasURL)
	if err != nil {
		return err
	}

	cfg.nextLocationAreasURL = resp.Next
	cfg.prevLocationAreasURL = resp.Previous

	for _, area := range resp.Results {
		fmt.Println(area.Name)
	}

	return nil
}
