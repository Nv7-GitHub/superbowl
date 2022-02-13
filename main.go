package main

import "fmt"

func main() {
	fmt.Println(GetTeamScore("lar/los_angeles_rams", 2, true))
	fmt.Println(GetTeamScore("cin/cincinnati-bengals", 2, true))

	/*
		Notes:
			- When losing is weighted same as winning (losing against bad team is weighted less that losing against good team), bengals win
			- When losing against bad team is weighted more, then rams win
			- Increasing recursion depth just makes rams win more
			- When losing against bad team is weighted less, then rams win more when recursion depth increases
			- Rams has average win score of 4.5, Bengals have average win score of 4.1, these teams are really close
	*/
	dataFile.Close()
}
