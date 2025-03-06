package cmds


import "fmt"

func commandHelp(config *Config, args ...string) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for _, cmd := range GetCommands() {
		fmt.Println(cmd.name, ":", cmd.description)
	}
	fmt.Println()
	return nil
}
