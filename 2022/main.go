package main

import (
	"fmt"
	"math"
)

const GameCount = 10

func main() {
	depth := 1
	ramsScore := GetTeamScore("lar/los_angeles_rams", depth)
	bengalsScore := GetTeamScore("cin/cincinnati-bengals", depth)

	diff := math.Abs(ramsScore - bengalsScore)
	diff = math.Pow(diff, 1/float64(depth+1))
	if ramsScore > bengalsScore {
		fmt.Printf("Rams win by %0.2f!\n", diff)
	} else {
		fmt.Printf("Bengals win by %0.2f!\n", diff)
	}

	/*
		Notes:
			- When losing is weighted same as winning (losing against bad team is weighted less that losing against good team), bengals win
			- When losing against bad team is weighted more, then rams win
			- Increasing recursion depth just makes rams win more
			- When losing against bad team is weighted less, then rams win more when recursion depth increases
			- Rams has average win score of 4.5, Bengals have average win score of 4.1, these teams are really close
			- When recursing to 2, rams win, but going to 3 Bengals win [so who is better?!?!?]
	*/
	dataFile.Close()
}
