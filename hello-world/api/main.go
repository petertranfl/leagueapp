package main

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
)

type SummonerRequest struct {
	SummonerName string `json:"SummonerName?"`
}

type SummonerData struct {
	profileIconId int
	name          string
	puuid         string
	summonerLevel int64
	revisionDate  int64
	id            string
	accountId     string
}

var myClient = &http.Client{Timeout: 10 * time.Second}

func getJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func getSummonerData(user SummonerRequest) (RiotResponse, error) {
	apiKey := os.Getenv("apiKey")

}

func main() {
	lambda.Start(getSummonerData)
}
