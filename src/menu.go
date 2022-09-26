package main

import (
	"encoding/json"
	"os"
)

type Menu []struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	PreparationTime  int    `json:"preparation-time"`
	Complexity       int    `json:"complexity"`
	CookingApparatus string `json:"cooking-apparatus"`
}

var menu Menu

func ParseMenu() {
	menuFile, _ := os.Open("menu.json")
	jsonParser := json.NewDecoder(menuFile)
	jsonParser.Decode(&menu)
}
