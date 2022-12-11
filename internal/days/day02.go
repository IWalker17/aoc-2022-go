package days

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func Day02() int {
	part1 := calculateYourScoreFromStrategy()
	log.Printf("result: %v", part1)
	return part1
}

type Player struct {
	Score int
}

type Choice struct {
	Name   string
	Points int
}

func calculateYourScoreFromStrategy() int {
	file, err := os.Open("internal/testdata/days/day02.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var opponentOptions = make(map[string]Choice)
	opponentOptions["A"] = Choice{Name: "Rock", Points: 1}
	opponentOptions["B"] = Choice{Name: "Paper", Points: 2}
	opponentOptions["C"] = Choice{Name: "Scissors", Points: 3}

	var yourOptions = make(map[string]Choice)
	yourOptions["X"] = Choice{Name: "Rock", Points: 1}
	yourOptions["Y"] = Choice{Name: "Paper", Points: 2}
	yourOptions["Z"] = Choice{Name: "Scissors", Points: 3}

	opponentScore := Player{Score: 0}
	yourScore := Player{Score: 0}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		playersChoices := strings.Split(line, " ")
		opponentsChoice := opponentOptions[playersChoices[0]]
		yourChoice := yourOptions[playersChoices[1]]

		if opponentsChoice.Name == yourChoice.Name {
			opponentScore.Score += 3
			yourScore.Score += 3
		}

		if opponentsChoice.Name == "Rock" && yourChoice.Name == "Paper" ||
			opponentsChoice.Name == "Paper" && yourChoice.Name == "Scissors" ||
			opponentsChoice.Name == "Scissors" && yourChoice.Name == "Rock" {
			yourScore.Score += 6 + yourChoice.Points
			opponentScore.Score += 0 + opponentsChoice.Points
			continue
		} else {
			yourScore.Score += 0 + yourChoice.Points
			opponentScore.Score += 6 + opponentsChoice.Points
		}
	}

	log.Printf("Your final score is: %v\n", yourScore.Score)
	log.Printf("The opponents final score is: %v\n", opponentScore.Score)
	return yourScore.Score
}
