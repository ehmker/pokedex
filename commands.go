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

	// fmt.Print("explore res: ",res)

	fmt.Printf("Exploring %s\n", area)
	fmt.Println("Found Pokemon: ")
	for _, pkm := range res.PokemonEncounters{
		fmt.Printf(" - %s\n", pkm.Pokemon.Name)
	}

	return nil
}