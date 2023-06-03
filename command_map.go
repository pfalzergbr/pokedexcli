package main

import (
	"fmt"
	"log"
)

func commandMap(cfg *config) error {
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreasURL)
	if err != nil {
		log.Fatal(err)
	}

	cfg.nextLocationAreasURL = resp.Next
	cfg.prevLocationAreasURL = resp.Previous

	for _, area := range resp.Results {
		fmt.Println(area.Name)
		fmt.Println(area.URL)
	}

	return nil
}