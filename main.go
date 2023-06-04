package main

import (
	"time"

	"github.com/pfalzergbr/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient pokeapi.Client
	nextLocationAreasURL *string
	prevLocationAreasURL *string
	caughtPokemon map[string]pokeapi.Pokemon
}


func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(time.Minute * 10),
		nextLocationAreasURL: nil,
		prevLocationAreasURL: nil,
		caughtPokemon: make(map[string]pokeapi.Pokemon),
	}
	startRepl(&cfg)
}
