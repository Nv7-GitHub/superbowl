package main

import (
	"fmt"
	"sort"
)

func handle(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	fetch()

	// Build map
	for _, game := range games {
		teams[game.Winner].Wins[game.Loser] = struct{}{}
		teams[game.Loser].Losses[game.Winner] = struct{}{}
	}

	// Assign team scores
	scores := make(map[string]int)
	for _, team := range teams {
		for _, otherTeam := range teams {
			if otherTeam.Abbrev == team.Abbrev {
				continue
			}

			if compare(team.Abbrev, otherTeam.Abbrev) > 0 {
				scores[team.Abbrev]++
			}
		}
	}

	// Make team list
	teamList := make([]string, 0, len(teams))
	for _, team := range teams {
		teamList = append(teamList, team.Abbrev)
	}

	sort.Slice(teamList, func(i, j int) bool { return scores[teamList[i]] > scores[teamList[j]] })

	for i, v := range teamList {
		fmt.Printf("%d. %s\n", i+1, teams[v].Name)
	}

	fmt.Println("Kansis City Chiefs beats Philadelphia Eagles by", compare("KC", "PHI"))
}
