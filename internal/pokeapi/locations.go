package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type PokeLocation struct {
	Name string `json:"name"`
	Url string `json:"url"`
}
type LocationResp struct{
	Count int `json:"count"`
	Next string `json:"next"`
	Previous string `json:"previous"`
	Results []PokeLocation `json:"results"`
}

func GetAllLocations(url string) LocationResp {
	json_body := checkResponse(url)
	var locations LocationResp
	err := json.Unmarshal(json_body, &locations)

	if err != nil {
		fmt.Println("error:", err)
		return LocationResp{}
	}
	return locations
}
func checkResponse(url string) []byte{  // ([]byte, error){
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

	return body

}
	


