package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Team struct {
	Name  string
	Games []Game
}

type Game struct {
	Score         int
	OpponentScore int
	Opponent      string
	ScoreMult     float64 // Makes postseason more important
}

var dataFile *os.File
var teams map[string]Team

func init() {
	// Check if exists
	_, err := os.Stat("data.json")
	if os.IsNotExist(err) {
		err = os.WriteFile("data.json", []byte("{}"), os.ModePerm)
		if err != nil {
			panic(err)
		}
	}

	dataFile, err = os.OpenFile("data.json", os.O_RDWR, 0)
	if err != nil {
		panic(err)
	}

	dec := json.NewDecoder(dataFile)
	err = dec.Decode(&teams)
	if err != nil {
		panic(err)
	}
}

func GetTeam(name string) Team {
	tm, exists := teams[name]
	if exists {
		return tm
	}

	// Fetch team
	fmt.Println("Fetching: ", name)
	resp, err := http.Get("https://www.espn.com/nfl/team/schedule/_/name/" + name) // lar/los_angeles_rams for LA Rams
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// Parse into JSON
	dat := string(body)
	dat = strings.Split(dat, "window['__espnfitt__']=")[1]
	dat = strings.Split(dat, ";</script>")[0]
	var data map[string]interface{}
	err = json.Unmarshal([]byte(dat), &data)
	if err != nil {
		panic(err)
	}

	// Parse JSON
	events := data["page"].(map[string]interface{})["content"].(map[string]interface{})["scheduleData"].(map[string]interface{})["teamSchedule"].([]interface{})
	games := make([]Game, 0)
	for _, ev := range events {
		scoreMult := float64(1)
		switch ev.(map[string]interface{})["title"].(string) {
		case "Preseason": // Not important
			scoreMult = 0

		case "Postseason": // Most important
			scoreMult = 2
		}

		event := ev.(map[string]interface{})["events"].(map[string]interface{})
		gameData := event["post"].([]interface{})
		pre := event["pre"].([]interface{})
		gameData = append(gameData, pre...) // They are of same type

		for _, game := range gameData {
			gm := game.(map[string]interface{})

			// Check if done
			if !(gm["status"].(map[string]interface{})["completed"].(bool)) {
				continue
			}
			if gm["customText"] == "BYE WEEK" {
				continue
			}

			// Get data
			opponent := gm["opponent"].(map[string]interface{})["links"].(string)
			result := gm["result"].(map[string]interface{})
			scoreV := result["currentTeamScore"].(string)
			opponentScoreV := result["opponentTeamScore"].(string)

			// Parse data
			score, err := strconv.Atoi(scoreV)
			if err != nil {
				panic(err)
			}
			opponentScore, err := strconv.Atoi(opponentScoreV)
			if err != nil {
				panic(err)
			}
			opponent = strings.TrimPrefix(opponent, "/nfl/team/_/name/")

			games = append(games, Game{
				Score:         score,
				OpponentScore: opponentScore,
				Opponent:      opponent,
				ScoreMult:     float64(scoreMult),
			})
		}
	}

	// Make team
	tm = Team{
		Name:  name,
		Games: games,
	}

	// Persist
	teams[name] = tm
	_, err = dataFile.Seek(0, 0)
	if err != nil {
		panic(err)
	}
	err = dataFile.Truncate(0)
	if err != nil {
		panic(err)
	}
	enc := json.NewEncoder(dataFile)
	err = enc.Encode(teams)
	if err != nil {
		panic(err)
	}

	return tm
}
