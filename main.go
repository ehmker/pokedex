package main

func main(){
	// fmt.Println(pokeapi.GetAllLocations().Results[0])
	startREPL()
}

type cliCommand struct {
	Name string
	Description string
	Config *Config
	Callback func(*Config) error
}

type Config struct{
	Next string
	Previous string
}

func getCommands() map[string]cliCommand {
	mapConfig := Config{}

	return map[string]cliCommand{
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    commandHelp,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    commandExit,
		},
		"map":{
			Name: "map",
			Description: "Displays the next 20 locations",
			Config: &mapConfig,
			Callback: commandMap,
		},
		"mapb": {
			Name: "mapb",
			Description: "Displays the previous 20 locations",
			Config: &mapConfig,
			Callback: commandMapBack,
		},
	}
}



