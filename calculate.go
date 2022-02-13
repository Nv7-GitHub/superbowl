package main

func GetTeamScore(teamName string, depth int) int {
	team := GetTeam(teamName)
	scores := make([]int, len(team.Games))
	for i, game := range team.Games {
		if depth <= 0 {
			scores[i] = game.Score - game.OpponentScore
		} else {
			scores[i] = (game.Score - game.OpponentScore) * GetTeamScore(game.Opponent, depth-1)
		}
	}

	average := 0
	for _, score := range scores {
		average += score
	}
	average /= len(scores)

	if average < 1 { // Clamp to 1 so that when used recursively it doesn't hurt the team to play a bad team
		return 1
	}
	return average
}
