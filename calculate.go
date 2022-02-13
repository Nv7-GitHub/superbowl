package main

func GetTeamScore(teamName string, depth int) float64 {
	if depth < 0 {
		return 1
	}

	team := GetTeam(teamName)

	// Get oppScores
	oppScores := make([]float64, len(team.Games))
	for i, game := range team.Games {
		oppScores[i] = GetTeamScore(game.Opponent, depth-1)
	}

	// Adjust oppScores so that lowest is 1
	min := oppScores[0]
	for _, score := range oppScores {
		if score < min {
			min = score
		}
	}
	for i := range oppScores {
		oppScores[i] += 1 - min
	}

	// Get max for later
	max := oppScores[0]
	for _, score := range oppScores {
		if score > max {
			max = score
		}
	}
	if depth == 0 {
		max++ // Make loses worth the loss when doing average diff (otherwise max - oppScores[i] would just be 0)
	}

	// Calculate scores
	scores := make([]float64, len(team.Games))
	for i, game := range team.Games {
		diff := float64(game.Score - game.OpponentScore)
		if diff > 0 { // Win, if you beat a good team it should add more
			scores[i] = diff * oppScores[i]
		} else { // Lose, if you lost to a bad team it should subtract more
			scores[i] = diff * (max - oppScores[i])
		}
		scores[i] *= game.ScoreMult
	}

	average := float64(0)
	for _, score := range scores {
		average += score
	}
	average /= float64(len(scores))

	return average
}
