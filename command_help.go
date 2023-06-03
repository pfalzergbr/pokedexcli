package main

import "fmt"

func commandHelp() error {
	fmt.Println("Welcome to the Gopherized Pokedex!")
	fmt.Println("")
	fmt.Println("Here's a list of commands:")
	fmt.Println("")

	// fmt.Println("- help: Show this help message")
	// fmt.Println("- exit: Close your cool console Pokedex")

	commandMap := getCommands()

	for _, command := range commandMap {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}

	return nil
}
