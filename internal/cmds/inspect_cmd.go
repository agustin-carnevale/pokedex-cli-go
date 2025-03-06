package cmds

import (
	"errors"
	"fmt"
)

func commandInspect(config *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	pokemonName := args[0]
	pokemon, ok := config.CaughtPokemon[pokemonName]
	if !ok {
		fmt.Println("you have not caught that pokemon")
	} else {
		fmt.Println("Name:", pokemon.Name)
		fmt.Println("Height:", pokemon.Height)
		fmt.Println("Weight:", pokemon.Weight)
		fmt.Println("Stats:")
		for _, s := range pokemon.Stats {
			fmt.Printf(" -%s: %v\n", s.Stat.Name, s.BaseStat)
		}
		fmt.Println("Types:")
		for _, t := range pokemon.Types {
			fmt.Printf(" -%s\n", t.Type.Name)
		}
	}
	return nil
}
