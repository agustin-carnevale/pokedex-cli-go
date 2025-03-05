package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/agustin-carnevale/pokedex/internal/cmds"
	"github.com/agustin-carnevale/pokedex/internal/pokeapi"
)

func main() {
	scanner := bufio.NewScanner((os.Stdin))
	pokeapi_client := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	config := cmds.Config{
		Next:           "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20",
		Previous:       "",
		PokeClient: pokeapi_client,
	}

	for {
		fmt.Print("Pokedex> ")
		if !scanner.Scan() {
			break
		}

		commands := cleanInput(scanner.Text())
		if len(commands) == 0 {
			continue
		}

		if command, exists := cmds.GetCommands()[commands[0]]; exists {
			param := ""
			if len(commands) > 1 {
				param = commands[1]
			}
			if err := command.Callback(&config, param); err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}

}

func cleanInput(text string) []string {
	words := strings.Fields(text)
	for i, word := range words {
		words[i] = strings.ToLower(word)
	}
	return words
}
