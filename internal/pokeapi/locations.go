package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type LocationResp struct{
	Count int 					`json:"count"`
	Next string 				`json:"next"`
	Previous string 			`json:"previous"`
	Results []struct {
					Name string `json:"name"`
					Url string 	`json:"url"`
					} 			`json:"results"`
}

func GetLocationsFromAPI(url string) (LocationResp, []byte) {
	json_body := checkResponse(url)
	var locations LocationResp
	err := json.Unmarshal(json_body, &locations)

	if err != nil {
		fmt.Println("error:", err)
		return LocationResp{}, json_body
	}
	return locations, json_body
}
func checkResponse(url string) []byte {  // ([]byte, error){
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
func GetLocationsFromCache(raw_json []byte) LocationResp {
	var locations LocationResp
	err := json.Unmarshal(raw_json, &locations)

	if err != nil {
		fmt.Println("error:", err)
		return LocationResp{}
	}
	return locations
}

