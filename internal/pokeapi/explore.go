package pokeapi

import (
	"encoding/json"
	"fmt"
)

type ExploreResp struct {
	PokemonEncounters []struct{
		Pokemon struct{
			Name string `json:"name"`
			Url string `json:"url"`
		} `json:"pokemon"`
		
	} `json:"pokemon_encounters"`
	// Pokemon []struct {
	// 	Name string `json:"name"`
	// 	Url string 	`json:"url"`
	// } 				`json:"pokemon"`
}


func GetExplorationFromAPI(url string) (ExploreResp, []byte) {
	json_body := checkResponse(url)
	var explore ExploreResp
	err := json.Unmarshal(json_body, &explore)

	fmt.Print("explore:", explore)

	if err != nil {
		fmt.Println("error:", err)
		return ExploreResp{}, json_body
	}
	return explore, json_body
}

func GetExplorationFromCache(raw_json []byte) ExploreResp {
	var explore ExploreResp
	err := json.Unmarshal(raw_json, &explore)

	if err != nil {
		fmt.Println("error:", err)
		return ExploreResp{}
	}
	return explore
}
