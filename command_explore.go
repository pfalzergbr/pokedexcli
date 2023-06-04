package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no location area provided")
	}

	la := args[0]

	LocationArea, err := cfg.pokeapiClient.GetLocationArea(la)
	if err != nil {
		return err
	}

	fmt.Printf("Pokemon in %s:\n", LocationArea.Name)
	fmt.Println("---------------------")

	for _, pokemon := range LocationArea.PokemonEncounters {
		fmt.Printf("- %s\n", pokemon.Pokemon.Name)
	}

	return nil
}
