package main

import (
	"math/rand"
	"errors"
	"fmt"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no pokemon provided")
	}

	pokemonName := args[0]

	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	const treshold = 50
	randNum := rand.Intn(pokemon.BaseExperience)
	if randNum > treshold {
		fmt.Printf("Failed to catch %s!\n", pokemonName)
		return nil
	}

	fmt.Printf("Caught %s!\n", pokemonName)
	
	// fmt.Printf("Pokemon in %s:\n", LocationArea.Name)
	// fmt.Println("---------------------")

	// for _, pokemon := range LocationArea.PokemonEncounters {
	// 	fmt.Printf("- %s\n", pokemon.Pokemon.Name)
	// }

	return nil
}
