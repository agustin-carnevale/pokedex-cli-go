package cmds


import (
	"fmt"
)

func commandMap(config *Config, args ...string) error {
	locationsResp, _ := config.PokeClient.GetLocations(config.Next)
	// update config obj with next and previous urls
	config.Next = locationsResp.Next
	config.Previous = locationsResp.Previous

	for _, location := range locationsResp.Results {
		fmt.Println(location.Name)
	}
	return nil
}
