package cmds

import "github.com/agustin-carnevale/pokedex/internal/pokeapi"

type Config struct {
	Next           string
	Previous       string
	PokeClient pokeapi.Client
}

type cliCommand struct {
	name        string
	description string
	Callback    func(config *Config, param string) error
}

func GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"map": {
			name:        "map",
			description: "List of locations",
			Callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "List of prev locations",
			Callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "List of Pokemons at Location",
			Callback:    commandExplore,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			Callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			Callback:    commandExit,
		},
	}
}
