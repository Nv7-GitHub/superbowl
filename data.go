package main

import "errors"

type Team struct {
	Name  string
	Games []Game
}

type Game struct {
	Score         int
	OpponentScore int
	Opponent      string
}

func GetTeam(name string) Team {
	panic(errors.New("not implemented"))
}
