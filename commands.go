package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

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

func commandExit() error {
	fmt.Println("Closing the pokedex... Bye!")
	os.Exit(0)
	return nil
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Show this help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Close your cool console Pokedex",
			callback:    commandExit,
		},
	}

}
