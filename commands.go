package main

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Show this help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "List 20 location areas, moving forwards",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "List 20 location areas, moving backwards",
			callback:    commandMapb,
		},
		"exit": {
			name:        "exit",
			description: "Close your cool console Pokedex",
			callback:    commandExit,
		},
	}

}
