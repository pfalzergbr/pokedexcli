package main

import (
	"fmt"
	"os"
)

func commandExit() error {
	fmt.Println("Closing the pokedex... Bye!")
	os.Exit(0)
	return nil
}
