package main

import "fmt"

func main() {
	fmt.Println(GetTeamScore("lar/los_angeles_rams", 3))
	fmt.Println(GetTeamScore("cin/cincinnati-bengals", 3))

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
