package cmds

import (
	"fmt"
	"os"
)

func commandExit(config *Config, param string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
