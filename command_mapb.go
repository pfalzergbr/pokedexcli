package main

import (
	"fmt"
	"log"
)

func commandMapb(cfg *config) error {
	if(cfg.prevLocationAreasURL == nil) {
		fmt.Println("No previous page!")
		return nil
	}
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.prevLocationAreasURL)
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