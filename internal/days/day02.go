package days

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func Day02() (int, int) {
	part1 := calculateScoreFromPartialStrategy()
	part2 := calculateScoreFromCompleteStrategy()
	return part1, part2
}

type Player struct {
	Score int
}

type Choice struct {
	Name   string
	Points int
}

func calculateScoreFromPartialStrategy() int {
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

	return yourScore.Score
}

func calculateScoreFromCompleteStrategy() int {
	file, err := os.Open("internal/testdata/days/day02.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var opponentOptions = make(map[string]Choice)
	opponentOptions["A"] = Choice{Name: "Rock", Points: 1}
	opponentOptions["B"] = Choice{Name: "Paper", Points: 2}
	opponentOptions["C"] = Choice{Name: "Scissors", Points: 3}

	var rockPaperScissorsPoints = make(map[string]int)
	rockPaperScissorsPoints["Rock"] = 1
	rockPaperScissorsPoints["Paper"] = 2
	rockPaperScissorsPoints["Scissors"] = 3

	var outcomeOptions = make(map[string]string)
	outcomeOptions["X"] = "Lose"
	outcomeOptions["Y"] = "Draw"
	outcomeOptions["Z"] = "Win"

	opponentScore := Player{Score: 0}
	yourScore := Player{Score: 0}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		oppsChoiceAndOutcome := strings.Split(line, " ")
		opponentsChoice := opponentOptions[oppsChoiceAndOutcome[0]]
		outcome := outcomeOptions[oppsChoiceAndOutcome[1]]

		if outcome == "Draw" {
			opponentScore.Score += 3 + opponentsChoice.Points
			yourScore.Score += 3 + opponentsChoice.Points
		}

		if outcome == "Lose" {
			if opponentsChoice.Name == "Rock" {
				opponentScore.Score += 6 + opponentsChoice.Points
				yourScore.Score += 0 + rockPaperScissorsPoints["Scissors"]
			}
			if opponentsChoice.Name == "Paper" {
				opponentScore.Score += 6 + opponentsChoice.Points
				yourScore.Score += 0 + rockPaperScissorsPoints["Rock"]
			}
			if opponentsChoice.Name == "Scissors" {
				opponentScore.Score += 6 + opponentsChoice.Points
				yourScore.Score += 0 + rockPaperScissorsPoints["Paper"]
			}
		}

		if outcome == "Win" {
			if opponentsChoice.Name == "Rock" {
				opponentScore.Score += 0 + opponentsChoice.Points
				yourScore.Score += 6 + rockPaperScissorsPoints["Paper"]
			}
			if opponentsChoice.Name == "Paper" {
				opponentScore.Score += 0 + opponentsChoice.Points
				yourScore.Score += 6 + rockPaperScissorsPoints["Scissors"]
			}
			if opponentsChoice.Name == "Scissors" {
				opponentScore.Score += 0 + opponentsChoice.Points
				yourScore.Score += 6 + rockPaperScissorsPoints["Rock"]
			}
		}
	}

	return yourScore.Score
}
