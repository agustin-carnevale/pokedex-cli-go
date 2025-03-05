package cmds

import (
	"fmt"
)

func commandExplore(config *Config, location string) error {
	pokemons, _ := config.PokeClient.GetPokemonsByLocation(location)

	for _, location := range pokemons {
		fmt.Println(location.Name)
	}
	return nil
}
