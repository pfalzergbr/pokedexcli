package main

import (
	"bufio"
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

func main() {
	fmt.Println("--=== Pokedex ===--")
	fmt.Println("Gopher edition")
	fmt.Println("")
	fmt.Println("Gotta catch 'em all!")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		scanner.Scan()

		text := scanner.Text()

		if scanner.Err() != nil {
			fmt.Println("Error reading input: ", scanner.Err())
			continue
		}

		if len(text) == 0 {
			continue
		}

		fmt.Println(text)

		command, ok := commandMap[text]
		if !ok {
			fmt.Println("Unknown command! Type 'help' for a list of commands.")
			continue
		}

		err := command.callback()
		if err != nil {
			fmt.Println("Error executing command: ", err)
			continue
		}

	}

}
