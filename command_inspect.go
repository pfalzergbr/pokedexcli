package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no pokemon provided")
	}

	pokemonName := args[0]

	pokemon, ok := cfg.caughtPokemon[pokemonName]

	if !ok {
		fmt.Printf("You haven't caught %s yet\n", pokemonName)
		return nil
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")

	for _, stat := range pokemon.Stats {
		fmt.Printf("- %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	for _, pokemonType := range pokemon.Types {
		fmt.Printf("- %s\n", pokemonType.Type.Name)
	}

	return nil
}
