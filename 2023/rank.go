package main

const depth = 1
const directWinScore = 3
const indirectWinScore = 1

func compare(a, b string) int {
	diff := 0

	// Direct win score
	if _, exists := teams[a].Wins[b]; exists {
		diff += directWinScore
	}
	if _, exists := teams[b].Wins[a]; exists {
		diff -= directWinScore
	}

	// Indirect win score
	winsA := getWinsRec(a, depth)
	winsB := getWinsRec(b, depth)
	lossesA := getLossRec(a, depth)
	lossesB := getLossRec(b, depth)

	// Calc wins
	for win := range winsA {
		if _, exists := lossesB[win]; exists {
			diff += indirectWinScore
		}
	}
	for win := range winsB {
		if _, exists := lossesA[win]; exists {
			diff -= indirectWinScore
		}
	}

	return diff
}

var winsRec = make(map[string]map[string]struct{})
var lossesRec = make(map[string]map[string]struct{})

func getWinsRec(team string, depth int) map[string]struct{} {
	_, exists := winsRec[team]
	if exists {
		return winsRec[team]
	} else {
		res := make(map[string]struct{})
		getWinsRecV(team, depth, res)
		winsRec[team] = res
		return res
	}
}

func getLossRec(team string, depth int) map[string]struct{} {
	_, exists := lossesRec[team]
	if exists {
		return lossesRec[team]
	} else {
		res := make(map[string]struct{})
		getLossRecV(team, depth, res)
		lossesRec[team] = res
		return res
	}
}

func getWinsRecV(team string, depth int, res map[string]struct{}) {
	for win := range teams[team].Wins {
		res[win] = struct{}{}
		if depth > 0 {
			getWinsRecV(win, depth-1, res)
		}
	}
}

func getLossRecV(team string, depth int, res map[string]struct{}) {
	for loss := range teams[team].Losses {
		res[loss] = struct{}{}
		if depth > 0 {
			getLossRecV(loss, depth-1, res)
		}
	}
}
