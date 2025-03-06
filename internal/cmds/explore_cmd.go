package cmds

import (
	"errors"
	"fmt"
)

func commandExplore(config *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}

	location := args[0]
	pokemons, _ := config.PokeClient.GetPokemonsByLocation(location)

	for _, location := range pokemons {
		fmt.Println(location.Name)
	}
	return nil
}
