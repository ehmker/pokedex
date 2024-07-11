package main

import (
	"time"

	"github.com/ehmker/pokedexcli/internal/pokecache"
)

func main(){
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
	Cache *pokecache.PokeCache
}

func getCommands() map[string]cliCommand {
	cache := pokecache.NewCache(5 * time.Minute)
	mapConfig := Config{
		Cache: &cache,
	}
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



