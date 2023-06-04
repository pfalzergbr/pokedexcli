package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {

	fmt.Println("")

	if len(cfg.caughtPokemon) == 0 {
		fmt.Println("Empty Pokedex. You haven't caught any pokemons yet")
		fmt.Println("")
		return nil
	}

	fmt.Println("Let's see...")
	fmt.Println("The list of pokemons you have caught so far:")
	for _, pokemon := range cfg.caughtPokemon {
		fmt.Printf("- %s\n", pokemon.Name)
	}
	fmt.Println("")

	return nil
}
