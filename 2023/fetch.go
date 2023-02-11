package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const weekCount = 18

type Team struct {
	Abbrev string
	Name   string
}

type Game struct {
	Winner string
	Loser  string
}

var teams = make(map[string]Team)
var games = make([]Game, 0)

func fetchWeek(num int) {
	fmt.Printf("Fetching week %d...\n", num)

	// Fetch
	url := fmt.Sprintf("https://www.espn.com/nfl/schedule/_/week/%d/year/2022/seasontype/2", num)
	res, err := http.Get(url)
	handle(err)
	body, err := io.ReadAll(res.Body)
	handle(err)

	// Parse
	part := strings.SplitN(strings.SplitN(string(body), "window['__espnfitt__']=", 2)[1], ";</script>", 2)[0]
	var data map[string]any
	err = json.Unmarshal([]byte(part), &data)
	handle(err)

	// Extract
	week := data["page"].(map[string]any)["content"].(map[string]any)["events"].(map[string]any)
	for _, day := range week {
		for _, event := range day.([]any) {
			if event.(map[string]any)["isTie"].(bool) || event.(map[string]any)["tbd"].(bool) {
				continue
			}

			competitors := event.(map[string]any)["competitors"].([]any)
			var winner string
			var loser string
			for _, team := range competitors {
				t := Team{
					Abbrev: team.(map[string]any)["abbrev"].(string),
					Name:   team.(map[string]any)["displayName"].(string),
				}
				teams[t.Abbrev] = t
				if v, ok := team.(map[string]any)["winner"]; ok && v.(bool) {
					winner = t.Abbrev
				} else {
					loser = t.Abbrev
				}
			}
			games = append(games, Game{
				Winner: winner,
				Loser:  loser,
			})
		}
	}
}

func fetch() {
	// Check for cache
	if _, err := os.Stat("data.json"); !os.IsNotExist(err) {
		fmt.Println("Loading cache...")
		file, err := os.Open("data.json")
		handle(err)
		defer file.Close()
		decoder := json.NewDecoder(file)

		// Decode data into teams and games
		var data map[string]any
		err = decoder.Decode(&data)
		handle(err)
		for k, v := range data["teams"].(map[string]any) {
			teams[k] = Team{
				Abbrev: v.(map[string]any)["Abbrev"].(string),
				Name:   v.(map[string]any)["Name"].(string),
			}
		}
		for _, v := range data["games"].([]any) {
			games = append(games, Game{
				Winner: v.(map[string]any)["Winner"].(string),
				Loser:  v.(map[string]any)["Loser"].(string),
			})
		}
	} else {
		for i := 1; i <= weekCount; i++ {
			fetchWeek(i)
		}
		fmt.Println("Saving to cache...")
		file, err := os.Create("data.json")
		handle(err)
		defer file.Close()
		encoder := json.NewEncoder(file)
		data := map[string]any{
			"teams": teams,
			"games": games,
		}
		err = encoder.Encode(data)
		handle(err)
	}
}
