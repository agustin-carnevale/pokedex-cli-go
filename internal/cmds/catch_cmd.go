package cmds

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(config *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	pokemonName := args[0]
	pokemon, err := config.PokeClient.GetPokemon(pokemonName)
	if err != nil {
		fmt.Printf("Uppps something went wrong, please check that %s is a valid pokemon!\n", pokemonName)
		return nil
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	catchDifficulty := rand.Intn(200)
	if catchDifficulty > pokemon.BaseExperience {
		fmt.Printf("%s was caught!\n", pokemonName)
		config.CaughtPokemon[pokemonName] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
	}

	return nil
}
