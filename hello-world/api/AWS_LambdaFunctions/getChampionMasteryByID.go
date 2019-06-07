package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
)

//ChampionMasteryRequest holds the data to use to lookup SummonerData
type ChampionMasteryRequest struct {
	SummonerName string `json:"SummonerName?"`
	SummonerID   string `json:"SummonerID?"`
}

//ChampionMasteryData holds the framework for Riot's Response
type ChampionMasteryData struct {
	ChampionLevel                int    `json:"championLevel,omitempty"`
	ChestGranted                 bool   `json:"chestGranted,omitempty"`
	ChampionPoints               int    `json:"championPoints,omitempty"`
	ChampionPointsSinceLastLevel int64  `json:"championPointsSinceLastLevel,omitempty"`
	ChampionPointsUntilNextLevel int64  `json:"championPointsUntilNextLevel, omitempty"`
	SummonerID                   string `json:summonerId, omitempty`
	TokensEarned                 int    `json:"tokensEarned,omitempty"`
	ChampionID                   int64  `json:"championId,omitempty"`
	LastPlayTime                 int64  `json:"lastPlayTime,omitempty"`
}

var myClient = &http.Client{Timeout: 10 * time.Second}

//getJSON reads JSON body and writes to r
func getJSON(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		log.Print("Error requesting data from Riot NA Server")
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

//getNASummonerDataByName requests SummonerData by name in NA
func getChampionMasteryByID(user ChampionMasteryRequest) (ChampionMasteryData, error) {
	apiKey := os.Getenv("apiKey")
	response := ChampionMasteryData{}
	err := getJSON("https://na1.api.riotgames.com/lol/champion-mastery/v4/champion-masteries/by-summoner/"+user.SummonerID+"?api_key="+apiKey, &response)
	if err != nil {
		log.Print("Error writing JSON to struct")
		log.Fatal(err)
	}
	return response, err
}

func main() {
	lambda.Start(getChampionMasteryByID)
}
