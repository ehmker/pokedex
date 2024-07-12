package main

import (
	"math/rand"
	"time"
)

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