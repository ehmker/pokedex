package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/ehmker/pokedexcli/internal/pokeapi"
)

func commandExit(c *Config) error {
	fmt.Println("Pokedex Exiting")
	os.Exit(0)
	return nil
}

func commandHelp(c *Config) error {
	fmt.Print("Usage:\n\n")
	for _, cmd := range getCommands(){
		fmt.Printf("%s: %s\n", cmd.Name, cmd.Description)
	}
	fmt.Println()
	return nil
}

func commandMap(c *Config) error {
	if c.Next == ""{
		c.Next = "https://pokeapi.co/api/v2/location-area/"
	}
	res := pokeapi.GetAllLocations(c.Next)
	
	c.Next = res.Next
	c.Previous = res.Previous

	for _, loc := range  res.Results{
		fmt.Println(loc.Name)
	}

	fmt.Println(c.Next, c.Previous)
	return nil
}

func commandMapBack(c *Config) error {
	if c.Previous == ""{
		return errors.New("unable to go back, already on first page")
	}
	res := pokeapi.GetAllLocations(c.Previous)

	c.Next = res.Next
	c.Previous = res.Previous

	for _, loc := range  res.Results{
		fmt.Println(loc.Name)
	}
	return nil
}
