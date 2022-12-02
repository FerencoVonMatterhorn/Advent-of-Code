package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	lines := []string{}

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var totalScore int
	for _, line := range lines {

		// part 1
		//roundScore, err := getScoreForRound(line[0:1], line[2:3])

		//part 2
		bestShape, err := getBestShape(line[0:1], line[2:3])
		if err != nil {
			log.Fatal(err)
		}

		roundScore, err := getScoreForRound(line[0:1], bestShape)
		if err != nil {
			log.Fatal(err)
		}
		totalScore += roundScore
	}

	fmt.Printf("Calculated points based on the table: %d", totalScore)

}

// A,X means rock
// B,Y means paper
// C,Z means scissors

func getScoreForRound(opponentShape string, myShape string) (int, error) {
	switch opponentShape {
	case "A":
		if myShape == "X" {
			return 4, nil
		} else if myShape == "Y" {
			return 8, nil
		} else if myShape == "Z" {
			return 3, nil
		} else {
			return 0, fmt.Errorf("myShape not recognized")
		}

	case "B":
		if myShape == "X" {
			return 1, nil
		} else if myShape == "Y" {
			return 5, nil
		} else if myShape == "Z" {
			return 9, nil
		} else {
			return 0, fmt.Errorf("myShape not recognized")
		}

	case "C":
		if myShape == "X" {
			return 7, nil
		} else if myShape == "Y" {
			return 2, nil
		} else if myShape == "Z" {
			return 6, nil
		} else {
			return 0, fmt.Errorf("myShape not recognized")
		}
	default:
		return 0, fmt.Errorf("opponentShape is not recognized")
	}
}

// getBestShape figures out what shape to choose for myself so the round ends as indicated.
// X means i need to lose
// Y means i need a draw
// Z means i need to win
func getBestShape(opponentShape string, roundOutcome string) (string, error) {
	switch opponentShape {
	case "A":
		if roundOutcome == "X" {
			return "Z", nil
		} else if roundOutcome == "Y" {
			return "X", nil
		} else if roundOutcome == "Z" {
			return "Y", nil
		} else {
			return "", fmt.Errorf("roundOutcome not recognized")
		}
	case "B":
		if roundOutcome == "X" {
			return "X", nil
		} else if roundOutcome == "Y" {
			return "Y", nil
		} else if roundOutcome == "Z" {
			return "Z", nil
		} else {
			return "", fmt.Errorf("roundOutcome not recognized")
		}
	case "C":
		if roundOutcome == "X" {
			return "Y", nil
		} else if roundOutcome == "Y" {
			return "Z", nil
		} else if roundOutcome == "Z" {
			return "X", nil
		} else {
			return "", fmt.Errorf("roundOutcome not recognized")
		}
	default:
		return "", fmt.Errorf("opponentShape not recognized")
	}
}
