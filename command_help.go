package main

import "fmt"

func commandHelp() error {
	fmt.Println("Welcome to the Gopherized Pokedex!")
	fmt.Println("")
	fmt.Println("Here's a list of commands:")
	fmt.Println("")
	
	commandMap := getCommands()

	for _, command := range commandMap {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}

	return nil
}
