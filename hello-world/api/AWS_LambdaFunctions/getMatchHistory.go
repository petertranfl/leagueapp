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

//MatchHistoryRequest holds the data to use to lookup Match History
type MatchHistoryRequest struct {
	AccountID string `json:"AccountID?"`
}

//MatchHistoryData holds the struct for Riot's Response
type MatchHistoryData struct {
	Matches    *[]MatchReference `json:"matches"`
	TotalGames int               `json:"totalGames"`
	StartIndex int               `json:"startIndex"`
	EndIndex   int               `json:"endIndex"`
}

//MatchReference stores the data for each single match record
type MatchReference struct {
	Lane       string `json:"lane"`
	GameID     int64  `json:"gameId"`
	Champion   int    `json:"champion"`
	PlatformID string `json:"platformId"`
	Season     int    `json:"season"`
	Queue      int    `json:"queue"`
	Role       string `json:"role"`
	TimeStamp  int64  `json:"timestamp"`
}

var myClient = &http.Client{Timeout: 10 * time.Second}

//getJSON reads JSON body and writes to target
func getJSON(url string, target *MatchHistoryData) error {
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
func getMatchHistory(user MatchHistoryRequest) (MatchHistoryData, error) {
	apiKey := os.Getenv("apiKey")
	var response MatchHistoryData
	err := getJSON("https://na1.api.riotgames.com/lol/match/v4/matchlists/by-account/"+user.AccountID+"?api_key="+apiKey, &response)
	if err != nil {
		log.Print("Error writing JSON to struct")
		log.Fatal(err)
	}
	return response, err
}

func main() {
	lambda.Start(getMatchHistory)
}
