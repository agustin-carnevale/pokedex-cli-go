package cmds

import (
	"fmt"
)

func commandMapb(config *Config, param string) error {
	if config.Previous == "" {
		fmt.Println("you're on the first page")
	} else {
		locationsResp, _ := config.PokeClient.GetLocations(config.Previous)
		// update config obj with next and previous urls
		config.Next = locationsResp.Next
		config.Previous = locationsResp.Previous

		for _, location := range locationsResp.Results {
			fmt.Println(location.Name)
		}
	}

	return nil
}
