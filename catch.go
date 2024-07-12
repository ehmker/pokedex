package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/ehmker/pokedexcli/internal/pokeapi"
)

type Pokemon struct{
	Name string
	Height int
	Weight int
	Stats map[string]int
	PokeType []string
}

// type output interface{
// 	PrintOutput()
// }

//base experience values range from 36-608
//catch will be successful if
//random n[0, 850] > baseExperience
func CatchPokemon(BaseExperience int) bool {
	baseTime := rand.Intn(5)
	waitTime := time.Duration(baseTime) * time.Second	
	time.Sleep(waitTime)
	n := rand.Intn(850)
	return n > BaseExperience
}

func AddToPokedex(pkmRsp pokeapi.PokemonResp, pokedex map[string]Pokemon) bool {
	
	if _, alreadyAdded := pokedex[pkmRsp.Name]; alreadyAdded {
		return false
	}

	statMap := make(map[string]int)
	var typeSlice []string
	for _, s := range pkmRsp.Stats{
		statMap[s.Stat.Name] = s.BaseStat
	}
	for _, t := range pkmRsp.Types{
		typeSlice = append(typeSlice, t.Type.Name)
	}
	pokedex[pkmRsp.Name] = Pokemon{
		Name: pkmRsp.Name,
		Height: pkmRsp.Height,
		Weight: pkmRsp.Weight,
		Stats: statMap,
		PokeType: typeSlice,
	}
	return true
}

func (pkm Pokemon) PrintOutput() {
	stat_order := []string{
		"hp", "attack", "defense", "special-attack", "special-defense", "speed",
	}
	fmt.Println("Name:", pkm.Name)
	fmt.Println("Height:", pkm.Height)
	fmt.Println("Weight:", pkm.Weight)
	fmt.Println("Stats:")
	for _, stat := range stat_order{
		fmt.Printf("  - %v: %v\n", stat, pkm.Stats[stat])
	} 
	fmt.Println("Types:")
	for _, t := range pkm.PokeType{
		fmt.Printf("  - %v\n", t)
	}
}