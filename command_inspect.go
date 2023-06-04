package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	fmt.Println("")
	if len(args) != 1 {
		return errors.New("no pokemon provided")
	}

	pokemonName := args[0]

	pokemon, ok := cfg.caughtPokemon[pokemonName]


	if !ok {
		fmt.Printf("You haven't caught %s yet\n", pokemon.Name)
		return nil
	}

	fmt.Println("Let's see...")
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)

	fmt.Println("Stats:")

	for _, stat := range pokemon.Stats {
		fmt.Printf("- %s: %d\n", (stat.Stat.Name), stat.BaseStat)
	}

	fmt.Println("Types:")
	for _, pokemonType := range pokemon.Types {
		fmt.Printf("- %s\n", (pokemonType.Type.Name))
	}
	fmt.Println("")
	return nil
}
