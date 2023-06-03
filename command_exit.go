package main

import (
	"fmt"
	"os"
)

func commandExit(cfg *config) error {
	fmt.Println("Closing the pokedex... Bye!")
	os.Exit(0)
	return nil
}
