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

//MatchDetailsRequest holds the data to use to lookup Match Details
type MatchDetailsRequest struct {
	MatchID string `json:"MatchID?"`
}

//MatchDetailsData holds the data for each match
type MatchDetailsData struct {
	SeasonID              int                          `json:"seasonId"`
	QueueID               int                          `json:"queueId"`
	GameID                int64                        `json:"gameId"`
	ParticipantIdentities []ParticipantIdentityDetails `json:"participantIdentities"`
	GameVersion           string                       `json:"gameVersion"`
	PlatformID            string                       `json:"platformId"`
	GameMode              string                       `json:"gameMode"`
	MapID                 int                          `json:"mapId"`
	GameType              string                       `json:"gameType"`
	Teams                 []TeamStats                  `json:"teams"`
	Participants          []ParticipantDetails         `json:"participants"`
	GameDuration          int64                        `json:"gameDuration"`
	GameCreation          int64                        `json:"gameCreation`
}

//ParticipantIdentityDetails links each player with their ParticipantID in each match
type ParticipantIdentityDetails struct {
	Player        PlayerDetails `json:"player"`
	ParticipantID int           `json:"participantId"`
}

//PlayerDetails holds the data for each Player in the match
type PlayerDetails struct {
	CurrentPlatformID string `json:"currentPlatformId"`
	SummonerName      string `json:"summonerName"`
	MatchHistoryURI   string `json:"matchHistoryUri"`
	PlatformID        string `json:"platformId"`
	CurrentAccountID  string `json:"currentAccountId`
	ProfileIcon       int    `json:"profileIcon"`
	SummonerID        string `json:"summonerId"`
	AccountID         string `json:"accountId"`
}

//TeamStats holds data for each team in a single match
type TeamStats struct {
	FirstDragon          bool              `json:"firstDragon"`
	FirstInhibitor       bool              `json:"firstInhibitor"`
	Bans                 []TeamBansDetails `json:"bans"`
	BaronKills           int               `json:"baronKills"`
	FirstRiftHerald      bool              `json:"firstRiftHerald"`
	FirstBaron           bool              `json:"firstBaron"`
	RiftHeraldKills      int               `json:"riftHeraldKills"`
	FirstBlood           bool              `json:"firstBlood"`
	TeamID               int               `json:"teamId"`
	FirstTower           bool              `json:"firstTower"`
	VileMawKills         int               `json:"vileMawKills,omitempty"`
	InhibitorKills       int               `json:"inhibitorKills"`
	TowerKills           int               `json:"towerKills"`
	DominionVictoryScore int               `json:"dominionVictoryScore,omitempty"`
	Win                  string            `json:"win"`
	DragonKills          int               `json:"dragonKills"`
}

//TeamBansDetails holds the team ban data in a single match
type TeamBansDetails struct {
	PickTurn   int `json:"pickTurn"`
	ChampionID int `json:"championId"`
}

//ParticipantDetails holds the data per player in each match
type ParticipantDetails struct {
	Stats                     ParticipantStatsDetails    `json:"stats"`
	ParticipantID             int                        `json:"participantId"`
	Runes                     RuneDetails                `json:"runes"`
	Timeline                  ParticipantTimelineDetails `json:"timeline"`
	TeamID                    int                        `json:"teamId"`
	Spell2Id                  int                        `json:"spell2Id"`
	Masteries                 MasteryDetails             `json:"masteries"`
	HighestAchievedSeasonTier string                     `json:"highestAchievedSeasonTier"`
	Spell1ID                  int                        `json:"spell1Id"`
	ChampionID                int                        `json:"championId"`
}

//ParticipantStatsDetails holds the stats for each player in each match
type ParticipantStatsDetails struct {
	FirstBloodAssist                bool  `json:"firstBloodAssist"`
	VisionScore                     int64 `json:"visionScore"`
	MagicDamageDealtToChampions     int64 `json:"magicDamageDealtToChampions`
	DamageDealtToObjectives         int64 `json:"damageDealtToObjectives"`
	TotalTimeCrowdControlDealt      int   `json:"totalTimeCrowdControlDealt"`
	LongestTimeSpentLiving          int   `json:"longestTimeSpentLiving"`
	Perk1Var1                       int   `json:"perk1Var"`
	Perk1Var3                       int   `json:"perk1Var3"`
	Perk1Var2                       int   `json:"perk1Var2"`
	TripleKills                     int   `json:"tripleKills"`
	Perk3Var3                       int   `json:"perk3Var3"`
	NodeNeutralizeAssist            int   `json:"nodeNeutralizeAssist"`
	Perk3Var2                       int   `json:"perk3Var2"`
	PlayerScore9                    int   `json:"playerScore9"`
	PlayerScore8                    int   `json:"playerScore8"`
	Kills                           int   `json:"kills"`
	PlayerScore1                    int   `json:"playerScore1"`
	PlayerScore0                    int   `json:"playerScore0"`
	PlayerScore3                    int   `json:"playerScore3"`
	PlayerScore2                    int   `json:"playerScore2"`
	PlayerScore5                    int   `json:"playerScore5"`
	PlayerScore4                    int   `json:"playerScore4"`
	PlayerScore7                    int   `json:"playerScore7"`
	PlayerScore6                    int   `json:"playerScore6"`
	Perk5Var1                       int   `json:"perk5Var1"`
	Perk5Var3                       int   `json:"perk5Var3"`
	Perk5Var2                       int   `json:"perk5Var2"`
	TotalScoreRank                  int   `json:"totalScoreRank"`
	NeutralMinionsKilled            int   `json:"neutralMinionsKilled"`
	DamageDealtToTurrets            int64 `json:"damageDealtToTurrets"`
	PhysicalDamageDealtToChampions  int64 `json:"physicalDamageDealtToChampions"`
	NodeCapture                     int   `json:"nodeCapture"`
	LargestMultiKill                int   `json:"largestMultiKill"`
	Perk2Var2                       int   `json:"perk2Var2"`
	Perk2Var3                       int   `json:"perk2Var3"`
	TotalUnitsHealed                int   `json:"totalUnitsHealed"`
	Perk2Var1                       int   `json:"perk2Var1"`
	Perk4Var1                       int   `json:"perk4Var1"`
	Perk4Var2                       int   `json:"perk4Var2"`
	Perk4Var3                       int   `json:"perk4Var3"`
	WardsKilled                     int   `json:"wardsKilled"`
	LargestCriticalStrike           int   `json:"largestCriticalStrike"`
	LargestKillingSpree             int   `json:"largestKillingSpree"`
	QuadraKills                     int   `json:"quadraKills"`
	TeamObjective                   int   `json:"teamObjective"`
	magicDamageDealt                int64 `json:"magicDamageDealt"`
	Item2                           int   `json:"item2"`
	Item3                           int   `json:"item3"`
	Item0                           int   `json:"item0"`
	NeutralMinionsKilledTeamJungle  int   `json:"neutralMinionsKilledTeamJungle"`
	Item6                           int   `json:"item6"`
	Item4                           int   `json:"item4"`
	Item5                           int   `json:"item5"`
	Perk1                           int   `json:"perk1"`
	Perk0                           int   `json:"perk0"`
	Perk3                           int   `json:"perk3"`
	Perk2                           int   `json:"Perk2"`
	Perk5                           int   `json:"Perk5"`
	Perk4                           int   `json:"Perk4"`
	Perk3Var1                       int   `json:"perk3Var1"`
	DamageSelfMitigated             int64 `json:"damageSelfMitigated"`
	MagicalDamageTaken              int64 `json:"magicalDamageTaken"`
	FirstInhibitorKill              bool  `json:"firstInhibitorKill"`
	TrueDamageTaken                 int64 `json:"trueDamageTaken"`
	NodeNeutralize                  int   `json:"nodeNeutralize"`
	Assists                         int   `json:"assists"`
	CombatPlayerScore               int   `json:"combatPlayerScore"`
	PerkPrimaryStyle                int   `json:"perkPrimaryStyle"`
	GoldSpent                       int   `json:"goldSpent"`
	TrueDamageDealt                 int64 `json:"trueDamageDealt"`
	ParticipantID                   int   `json:"participantId"`
	TotalDamageTaken                int64 `json:"totalDamageTaken"`
	PhysicalDamageDealt             int64 `json:"physicalDamageDealt"`
	SightWardsBoughtInGame          int   `json:"sightWardsBoughtInGame"`
	TotalDamageDealtToChampions     int64 `json:"totalDamageDealtToChampions"`
	PhysicalDamageTaken             int64 `json:"physicalDamageTaken"`
	TotalPlayerScore                int   `json:"totalPlayerScore"`
	Win                             bool  `json:"win"`
	ObjectivePlayerScore            int   `json:"objectivePlayerScore"`
	TotalDamageDealt                int64 `json:"totalDamageDealt"`
	Item1                           int   `json:"item1"`
	NeutralMinionsKilledEnemyJungle int   `json:"neutralMinionsKilledEnemyJungle"`
	Deaths                          int   `json:"deaths"`
	WardsPlaced                     int   `json:"wardsPlaced"`
	PerkSubStyle                    int   `json:"perkSubStyle"`
	TurretKills                     int   `json:"turretKills"`
	FirstBloodKill                  bool  `json:"firstBloodKill"`
	TrueDamageDealtToChampions      int64 `json:"trueDamageDealtToChampions"`
	GoldEarned                      int   `json:"goldEarned"`
	KillingSprees                   int   `json:"killingSprees"`
	UnrealKills                     int   `json:"unrealKills"`
	AltarsCaptured                  int   `json:"altarsCaptured"`
	FirstTowerAssist                bool  `json:"firstTowerAssist"`
	FirstTowerKill                  bool  `json:"firstTowerKill"`
	ChampLevel                      int   `json:"champLevel"`
	DoubleKills                     int   `json:"doubleKills"`
	NodeCaptureAssist               int   `json:"nodeCaptureAssist"`
	InhibitorKills                  int   `json:"inhibitorKills"`
	FirstInhibitorAssist            bool  `json:"firstInhibitorAssist"`
	Perk0Var1                       int   `json:"perk0Var1"`
	Perk0Var2                       int   `json:"perk0Var2"`
	Perk0Var3                       int   `json:"perk0Var3"`
	VisionWardsBoughtInGame         int   `json:"visionWardsBoughtInGame"`
	AltarsNeutralized               int   `json:"altarsNeutralized"`
	PentaKills                      int   `json:"pentaKills"`
	TotalHeal                       int64 `json:"totalHeal"`
	TotalMinionsKilled              int   `json:"totalMinionsKilled"`
	TimeCCingOthers                 int64 `json:"timeCCingOthers"`
}

//RuneDetails holds the rune data per player
type RuneDetails struct {
	RuneID int `json:"runeId"`
	Rank   int `json:"rank"`
}

//ParticipantTimelineDetails holds the timeline data per player per match
type ParticipantTimelineDetails struct {
	Lane                        string             `json:"lane"`
	ParticipantID               int                `json:"participantId"`
	CSDiffPerMinDeltas          map[string]float64 `json:"csDiffPerMinDeltas"`
	GoldPerPerMinDeltas         map[string]float64 `json:"goldPerMinDeltas"`
	XPDiffPerMinDeltas          map[string]float64 `json:"xpDiffPerMinDeltas"`
	CreepsPerMinDeltas          map[string]float64 `json:"creepsPerMinDeltas"`
	XPPerMinDeltas              map[string]float64 `json:"xpPerMinDeltas"`
	Role                        string             `json:"role"`
	DamageTakenDiffPerMinDeltas map[string]float64 `json:"damageTakenDiffPerMinDeltas"`
	DamageTakenPerMinDeltas     map[string]float64 `json:"damageTakenPerMinDeltas"`
}

//MasteryDetails holds the mastery data per player per match
type MasteryDetails struct {
	MasteryID int `json:"masteryId"`
	Rank      int `json:"rank"`
}

var myClient = &http.Client{Timeout: 10 * time.Second}

//getJSON reads JSON body and writes to target
func getJSON(url string, target *MatchDetailsData) error {
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
func getMatchDetails(user MatchDetailsRequest) (MatchDetailsData, error) {
	apiKey := os.Getenv("apiKey")
	var response MatchDetailsData
	err := getJSON("https://na1.api.riotgames.com/lol/match/v4/matches/"+user.MatchID+"?api_key="+apiKey, &response)
	if err != nil {
		log.Print("Error writing JSON to struct")
		log.Fatal(err)
	}
	return response, err
}

func main() {
	lambda.Start(getMatchDetails)
}
