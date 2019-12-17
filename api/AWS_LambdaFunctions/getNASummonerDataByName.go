package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
)

//SummonerRequest holds the data to use to lookup SummonerData
type SummonerRequest struct {
	SummonerName string `json:"SummonerName?"`
	AccountID    string `json:"AccountID?"`
}

//SummonerData holds the framework for Riot's Response
type SummonerData struct {
	ProfileIconID int    `json:"profileIconId,omitempty"`
	Name          string `json:"name,omitempty"`
	PuuID         string `json:"puuid,omitempty"`
	SummonerLevel int64  `json:"summonerLevel,omitempty"`
	RevisionDate  int64  `json:"revisionDate,omitempty"`
	SummonerID    string `json:"id,omitempty"`
	AccountID     string `json:"accountId,omitempty"`
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
func getNASummonerDataByName(user SummonerRequest) (SummonerData, error) {
	apiKey := os.Getenv("apiKey")
	response := SummonerData{}
	err := getJSON("https://na1.api.riotgames.com/lol/summoner/v4/summoners/by-name/"+user.SummonerName+"?api_key="+apiKey, &response)
	if err != nil {
		log.Print("Error writing JSON to struct")
		log.Fatal(err)
	}
	return response, err
}

func main() {
	lambda.Start(getNASummonerDataByName)
}
