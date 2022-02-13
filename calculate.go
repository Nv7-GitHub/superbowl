package main

func GetTeamScore(teamName string, depth int, noclamp ...bool) float64 {
	if depth < 0 {
		return 1
	}

	team := GetTeam(teamName)

	// Get oppScores
	oppScores := make([]float64, len(team.Games))
	for i, game := range team.Games {
		oppScores[i] = GetTeamScore(game.Opponent, depth-1)
	}
	max := float64(-1)
	for _, score := range oppScores {
		if score > max {
			max = score
		}
	}
	if depth == 0 {
		max++ // Make loses worth the loss when doing average diff (otherwise max - oppScores[i] would just be 0)
	}

	scores := make([]float64, len(team.Games))
	for i, game := range team.Games {
		diff := float64(game.Score - game.OpponentScore)
		if diff > 0 { // Win, if you beat a good team it should add more
			scores[i] = diff * oppScores[i]
		} else { // Lose, if you lost to a bad team it should subtract more
			scores[i] = diff * (max - oppScores[i])
		}
	}

	average := float64(0)
	for _, score := range scores {
		average += score
	}
	average /= float64(len(scores))

	if average < 1 && len(noclamp) == 0 { // Clamp to 1 so that when used recursively it doesn't hurt the team to play a bad team
		return 1
	}
	return average
}
