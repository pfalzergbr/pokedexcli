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

func cleanInput(str string) []string {
	lower := strings.ToLower(str)
	tokens := strings.Fields(lower)
	return tokens
}