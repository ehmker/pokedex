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
	Callback func(*Config, string) error
}

type Config struct{
	Next string
	Previous string
	Cache *pokecache.PokeCache
	Pokedex map[string]Pokemon
}

func getCommands() map[string]cliCommand {
	cache := pokecache.NewCache(5 * time.Minute)
	pokedex := make(map[string]Pokemon)
	config := Config{
		Cache: &cache,
		Pokedex: pokedex,
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
			Description: "Displays the next 20 locations.",
			Config: &config,
			Callback: commandMap,
		},
		"mapb": {
			Name: "mapb",
			Description: "Displays the previous 20 locations.",
			Config: &config,
			Callback: commandMapBack,
		},
		"explore": {
			Name: "explore",
			Description: "Called as 'explore <location name/id>'.\n\tDisplays the pokemon that appear in area.",
			Config: &config,
			Callback: commandExplore,
		},
		"catch": {
			Name: "catch",
			Description: "Called as 'catch <pokemon name/id>'.\n\tAttempts to catch the pokemon and add to the pokedex.",
			Config: &config,
			Callback: commandCatch,
		},
		"inspect":{
			Name: "inspect",
			Description: "Called as 'inspect <pokemon name/id>'.\n\tPrints the base stats of the given pokemon if it has already been caught.",
			Config: &config,
			Callback: commandInspect,
		},
		"pokedex":{
			Name: "pokedex",
			Description: "Displays the pokemon within the PokeDex",
			Config: &config,
			Callback: commandPokedex,
		},
	}
}