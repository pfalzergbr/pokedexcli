package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
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

	fmt.Println("Pokeball, go!")
	fmt.Println("Hit!")

	for i := 0; i < 3; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println("wiggle - wiggle...")
	}

	const treshold = 50
	randNum := rand.Intn(pokemon.BaseExperience)
	if randNum > treshold {
		fmt.Printf("%s breaks free!\n", pokemonName)
		fmt.Printf("Failed to catch %s!\n", pokemonName)
		return nil
	}
	// Add to the map if we caught it. 
	cfg.caughtPokemon[pokemonName] = pokemon
	
	fmt.Println("Gotcha!")
	fmt.Printf("Caught %s!\n", pokemonName)

	// fmt.Printf("Pokemon in %s:\n", LocationArea.Name)
	// fmt.Println("---------------------")

	// for _, pokemon := range LocationArea.PokemonEncounters {
	// 	fmt.Printf("- %s\n", pokemon.Pokemon.Name)
	// }

	return nil
}
