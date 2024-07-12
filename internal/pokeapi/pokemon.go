package pokeapi

import (
	"encoding/json"
	"fmt"
)

type PokemonResp struct {
	BaseExperience 			int `json:"base_experience"`
	Height    				int `json:"height"`
	ID                     	int    `json:"id"`
	Name          			string `json:"name"`
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
	Weight int `json:"weight"`
}

func GetPokemonFromAPI(url string) (PokemonResp, []byte) {
	json_body := checkResponse(url)
	var pokemon PokemonResp
	err := json.Unmarshal(json_body, &pokemon)

	if err != nil {
		fmt.Println("error:", err)
		return PokemonResp{}, json_body
	}
	return pokemon, json_body
}



func GetPokemonFromCache(raw_json []byte) PokemonResp {
	var pokemon PokemonResp
	err := json.Unmarshal(raw_json, &pokemon)

	if err != nil {
		fmt.Println("error:", err)
		return PokemonResp{}
	}
	return pokemon
}