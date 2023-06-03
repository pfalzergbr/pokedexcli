package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
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

		parsedCommand := parseInput(text)
		
		if len(parsedCommand) == 0 {
			continue
		}

		command, ok := commandMap[parsedCommand[0]]
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

func parseInput(str string) []string {
	lower := strings.ToLower(str)
	tokens := strings.Fields(lower)
	return tokens
}
