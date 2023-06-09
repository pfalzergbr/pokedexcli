package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {
	fmt.Println("--=== Pokedex ===--")
	fmt.Println("Gopher edition")
	fmt.Println("")
	fmt.Println("Gotta catch 'em all!")

	scanner := bufio.NewScanner(os.Stdin)
	commandMap := getCommands()

	for {
		fmt.Print("Pokedex > ")
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

		command := parsedCommand[0]

		args := []string{}

		if len(parsedCommand) > 1{	
			args = parsedCommand[1:]
		}

		c, ok := commandMap[command]
		if !ok {
			fmt.Println("Unknown command! Type 'help' for a list of commands.")
			continue
		}

		err := c.callback(cfg, args...)
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
