package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
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
}

var dataFile *os.File
var teams map[string]Team

func init() {
	var err error
	dataFile, err = os.Open("data.json")
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
	schedule := data["page"].(map[string]interface{})["content"].(map[string]interface{})["scheduleData"].(map[string]interface{})["teamSchedule"].([]interface{})
	fmt.Println(schedule)

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
