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

	fmt.Println("help: Show this help message")
	fmt.Println("exit: Close your cool console Pokedex")
	// for _, command := range commandMap {
	// 	fmt.Printf("%s: %s", command.name, command.description)
	// }

	return nil
}

func commandExit() error {
	fmt.Println("Closing the pokedex... Bye!")
	os.Exit(0)
	return nil
}

var commandMap = map[string]cliCommand{
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
