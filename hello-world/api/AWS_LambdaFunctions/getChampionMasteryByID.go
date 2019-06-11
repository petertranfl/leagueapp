package main

import (
	"encoding/json"
	"io/ioutil"
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

//MasteryDataArray holds the struct for Riot's Response
type MasteryDataArray []struct {
	ChampionMasteryData []struct {
		ChestGranted                 bool   `json:"chestGranted,omitempty"`
		ChampionLevel                int    `json:"championLevel,omitempty"`
		ChampionPoints               int    `json:"championPoints,omitempty"`
		ChampionID                   int64  `json:"championId,omitempty"`
		ChampionPointsUntilNextLevel int64  `json:"championPointsUntilNextLevel,omitempty"`
		LastPlayTime                 int64  `json:"lastPlayTime,omitempty"`
		TokensEarned                 int    `json:"tokensEarned,omitempty"`
		ChampionPointsSinceLastLevel int64  `json:"championPointsSinceLastLevel,omitempty"`
		SummonerID                   string `json:summonerId,omitempty`
	}
}

var myClient = &http.Client{Timeout: 10 * time.Second}

//getJSON reads JSON body and writes to target
func getJSON(url string, target MasteryDataArray) error {
	r, err := myClient.Get(url)
	if err != nil {
		log.Print("Error requesting data from Riot NA Server")
		return err
	}
	defer r.Body.Close()

	res, err := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(res, &target)
	return err
}

//getChampionMasteryByID requests ChampionMastery by SummonerID in NA
func getChampionMasteryByID(user ChampionMasteryRequest) (MasteryDataArray, error) {
	apiKey := os.Getenv("apiKey")
	response := MasteryDataArray{}
	err := getJSON("https://na1.api.riotgames.com/lol/champion-mastery/v4/champion-masteries/by-summoner/"+user.SummonerID+"?api_key="+apiKey, response)
	if err != nil {
		log.Print("Error writing JSON to struct")
		log.Fatal(err)
	}
	return response, err
}

func main() {
	lambda.Start(getChampionMasteryByID)
}
