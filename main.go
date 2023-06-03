package main

import 	"github.com/pfalzergbr/pokedexcli/internal/pokeapi"

type config struct {
	pokeapiClient pokeapi.Client
	nextLocationAreasURL *string
	prevLocationAreasURL *string
}


func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(),
		nextLocationAreasURL: nil,
		prevLocationAreasURL: nil,

	}
	startRepl(&cfg)
}
