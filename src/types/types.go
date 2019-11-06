package main

import (
	"fmt"
	"time"
)

type Score struct {
	scoreInt int
}

type FootballPlayer struct {
	Name     string
	BirthDay time.Time
}

type FootballTeam []FootballPlayer

func (score *Score) Increment() {
	score.scoreInt++
}

func NewScore(value int) Score {
	return Score{scoreInt: value}
}

func (score Score) String() string {
	return fmt.Sprintf("%v", score.scoreInt)
}

func main() {
	score := NewScore(0)
	score.Increment()
	fmt.Println(score)
}
