package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Guild struct {
	Players []Player `json:"players"`
}
type Player struct {
	Units []Unit     `json:"units"`
	Data  PlayerData `json:"data"`
}

type PlayerData struct {
	Name string `json:"name"`
}

type Unit struct {
	Data DataType `json:"data"`
}

type DataType struct {
	BaseId      string    `json:"base_id"`
	GearLevel   int       `json:"gear_level"`
	AbilityData []Ability `json:"ability_data"`
	Zetas       []string  `json:"zeta_abilities"`
	Power       int       `json:"power"`
}

type Ability struct {
	Name    string `json:"name"`
	IsOmega bool   `json:"is_omega"`
	IsZeta  bool   `json:"is_zeta"`
}

func main() {
	fmt.Println()
	fmt.Println()
	// response, err := http.Get("https://swgoh.gg/api/guild/450/?format=json")
	//
	// if err != nil {
	// 	fmt.Print(err.Error())
	// 	os.Exit(1)
	// }
	//
	// responseData, err := ioutil.ReadAll(response.Body)
	// Open our jsonFile
	jsonFile, err := os.Open("lwoti.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var guild Guild
	const toonID = "HERMITYODA"
	json.Unmarshal(byteValue, &guild)
	fmt.Println(toonID)
	for _, guildie := range guild.Players {
		for _, toon := range guildie.Units {
			if toon.Data.BaseId == toonID {
				if toon.Data.GearLevel < 10 {
					continue
				}
				fmt.Printf("%-25s  ", guildie.Data.Name)
				fmt.Printf("G%-2d  -- %d -- [ ", toon.Data.GearLevel, toon.Data.Power)
				for _, zeta := range toon.Data.Zetas {
					fmt.Printf("%s ", zeta)
				}
				fmt.Println("]")
			}
		}
	}

}
