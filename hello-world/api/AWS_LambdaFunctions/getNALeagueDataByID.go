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

//LeagueDataRequest holds the data to use to lookup SummonerData
type LeagueDataRequest struct {
	SummonerID string `json:"SummonerID?"`
}

//NALeagueData holds the struct for Riot's Response
type NALeagueData struct {
	QueueType    string        `json:"queueType,omitempty"`
	SummonerName string        `json:"summonerName,omitempty"`
	HotStreak    bool          `json:"hotStreak,omitempty"`
	MiniSeries   MiniSeriesDTO `json:"miniSeries,omitempty"`
	Wins         int           `json:"wins,omitempty"`
	Veteran      bool          `json:"veteran,omitempty"`
	Losses       int           `json:"losses,omitempty"`
	Rank         string        `json:"rank,omitempty"`
	LeagueID     string        `json:"leagueId,omitempty"`
	Inactive     bool          `json:"inactive,omitempty"`
	FreshBlood   bool          `json:"freshBlood,omitempty"`
	Tier         string        `json:"tier,omitempty"`
	SummonerID   string        `json:"summonerId,omitempty"`
	LeaguePoints int           `json:"leaguePoints,omitempty"`
}

//MiniSeriesDTO holds the data for the promo games
type MiniSeriesDTO struct {
	Progress string `json:"progress,omitempty"`
	Losses   int    `json:"losses,omitempty"`
	Target   int    `json:"target,omitempty"`
	Wins     int    `json:"wins,omitempty"`
}

var myClient = &http.Client{Timeout: 10 * time.Second}

//getJSON reads JSON body and writes to target
func getJSON(url string, target *[]NALeagueData) error {
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
func getNALeagueDataByID(user LeagueDataRequest) ([]NALeagueData, error) {
	apiKey := os.Getenv("apiKey")
	var response []NALeagueData
	err := getJSON("https://na1.api.riotgames.com/lol/league/v4/entries/by-summoner/"+user.SummonerID+"?api_key="+apiKey, &response)
	if err != nil {
		log.Print("Error writing JSON to struct")
		log.Fatal(err)
	}
	return response, err
}

func main() {
	lambda.Start(getNALeagueDataByID)
}
