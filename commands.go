package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/ehmker/pokedexcli/internal/pokeapi"
)

func commandExit(c *Config, _ string) error {
	fmt.Println("Pokedex Exiting")
	os.Exit(0)
	return nil
}

func commandHelp(c *Config, _ string) error {
	fmt.Print("Usage:\n\n")
	for _, cmd := range getCommands(){
		fmt.Printf("%s: %s\n", cmd.Name, cmd.Description)
	}
	fmt.Println()
	return nil
}

func commandMap(c *Config, _ string) error {
	if c.Next == ""{
		if c.Previous != ""{
			return errors.New("unable to go forward, already on last page")
		}
		c.Next = "https://pokeapi.co/api/v2/location-area/"
	}
	res := pokeapi.LocationResp{}
	if cached_json, ok := c.Cache.Get(c.Next); ok {
		res = pokeapi.GetLocationsFromCache(cached_json)
	} else {
		var raw_json []byte
		res, raw_json = pokeapi.GetLocationsFromAPI(c.Next)
		c.Cache.Add(c.Next, raw_json)
	}
	
	
	c.Next = res.Next
	c.Previous = res.Previous

	for _, loc := range  res.Results{
		fmt.Println(loc.Name)
	}

	fmt.Println(c.Next, c.Previous)
	return nil
}

func commandMapBack(c *Config, _ string) error {
	if c.Previous == ""{
		return errors.New("unable to go back, already on first page")
	}

	res := pokeapi.LocationResp{}
	if cached_json, ok := c.Cache.Get(c.Previous); ok {
		res = pokeapi.GetLocationsFromCache(cached_json)
	} else {
		var raw_json []byte
		res, raw_json = pokeapi.GetLocationsFromAPI(c.Previous)
		c.Cache.Add(c.Previous, raw_json)
	}


	c.Next = res.Next
	c.Previous = res.Previous

	for _, loc := range  res.Results{
		fmt.Println(loc.Name)
	}
	return nil
}

func commandExplore(c *Config, area string) error {
	if area == ""{
		return errors.New("unable to explore. no area given")
	} 
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s/", area)
	res := pokeapi.ExploreResp{}
	
	if cached_json, ok := c.Cache.Get(url); ok{
		res = pokeapi.GetExplorationFromCache(cached_json)
	} else {
		var raw_json []byte
		res, raw_json = pokeapi.GetExplorationFromAPI(url)
		c.Cache.Add(url, raw_json)
	}



	fmt.Printf("Exploring %s\n", area)
	fmt.Println("Found Pokemon: ")
	for _, pkm := range res.PokemonEncounters{
		fmt.Printf(" - %s\n", pkm.Pokemon.Name)
	}

	return nil
}

func commandCatch(c *Config, pkm string) error {
	if pkm == ""{
		return errors.New("unable to catch. no pokemon given for attempt")
	}

	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s/", pkm)
	res := pokeapi.PokemonResp{}

	if cached_json, ok := c.Cache.Get(url); ok {
		res = pokeapi.GetPokemonFromCache(cached_json)
	} else {
		var raw_json []byte
		res, raw_json = pokeapi.GetPokemonFromAPI(url)
		c.Cache.Add(url, raw_json)
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", res.Name)
	if !CatchPokemon(res.BaseExperience){
		fmt.Printf("%s escaped!\n", res.Name)
		return nil
	}
	fmt.Printf("%s was caught!\n", res.Name)
	wasAdded := AddToPokedex(res, c.Pokedex)
	if wasAdded{
		fmt.Printf("%s was added to the PokeDex!\n\tYou may now inspect it with the inspect command", res.Name)
	}
	return nil
}

func commandInspect(c *Config, pkm string) error {
	if pkm == ""{
		return errors.New("unable to inspect. no pokemon given")
	}

	if p, ok := c.Pokedex[pkm]; !ok {
		fmt.Printf("'%s' has not been caught.\n  Use the 'catch' command to attempt to catch it\n", pkm)
		return errors.New("pokemon not found")
	} else{
		p.PrintOutput()
		return nil
	}
	
}

func commandPokedex(c *Config, _ string) error {
	fmt.Println("Your PokeDex:")
	for key := range c.Pokedex{
		fmt.Printf(" - %s\n", key)
	}
	return nil
}