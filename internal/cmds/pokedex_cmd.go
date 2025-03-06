package cmds

import (
	"fmt"
)

func commandPokedex(config *Config, args ...string) error {
	if len(config.CaughtPokemon) < 1 {
		fmt.Println("Your pokedex is empty")
	} else {
		fmt.Println("Your Pokedex:")
		for pokemonName, _ := range config.CaughtPokemon {
			fmt.Printf(" - %s\n", pokemonName)
		}
	}
	return nil
}
