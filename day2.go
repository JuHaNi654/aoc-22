package main

import (
	"fmt"
	"os"
	"strings"
)

var (
	scores = map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
		"X": 1,
		"Y": 2,
		"Z": 3,
	}
)

func runMatch(p1 string, p2 string) int {
	score1 := scores[p1]
	score2 := scores[p2]

	sum := score1 + score2

	if score1 == score2 {
		return score2 + 3
	}

	switch sum ^ score2 {
	case 1, 5, 6:
		return score2 + 6
	default:
		return score2
	}
}

func runMatch2(p1 string, p2 string) int {
	score1 := scores[p1]

	switch p2 {
	case "Z":
		values := [3]int{1, 2, 3}
		return 6 + values[score1%3]
	case "X":
		values := [3]int{2, 3, 1}
		return 0 + values[score1%3]
	default:
		return 3 + score1
	}
}

func main() {
	file, err := os.ReadFile("./file.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	matches := strings.Split(string(file), "\n")
	var match1 int
	for _, match := range matches {
		if match != "" {
			players := strings.Split(match, " ")
			match1 += runMatch(players[0], players[1])
		}
	}

	fmt.Printf("Part 1: %d\n\n", match1)

	var match2 int
	for _, match := range matches {
		if match != "" {
			players := strings.Split(match, " ")
			match2 += runMatch2(players[0], players[1])
		}
	}

	fmt.Printf("Part 2: %d\n\n", match2)

}
