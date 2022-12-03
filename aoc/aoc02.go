package main

import (
	"log"
	"strings"
)

var points = map[string]int{
	"X": 1,
	"Y": 2,
	"Z": 3,
}
var winMoves = map[string]string{
	"A": "Y",
	"B": "Z",
	"C": "X",
}
var drawMoves = map[string]string{
	"A": "X",
	"B": "Y",
	"C": "Z",
}
var loseMoves = map[string]string{
	"A": "Z",
	"B": "X",
	"C": "Y",
}

const winPoint = 6
const drawPoint = 3

// X means you need to lose, Y means you need to end the round in a draw, and Z means you need to win
func rockPaperScissors2() {
	scanner := scanner("resources/aoc_02.txt")
	score := 0
	for scanner.Scan() {
		text := scanner.Text()
		game := strings.Split(text, " ")
		opp := game[0]
		mine := game[1]

		if mine == "Y" { // draw
			score = score + drawPoint
			score = score + points[drawMoves[opp]]
		} else if mine == "Z" { // win
			score = score + winPoint
			score = score + points[winMoves[opp]]
		} else { // lose
			score = score + points[loseMoves[opp]]
		}
	}
	log.Printf("Score %v", score)
}

func rockPaperScissors1() {
	scanner := scanner("resources/aoc_02.txt")
	score := 0
	for scanner.Scan() {
		text := scanner.Text()
		game := strings.Split(text, " ")
		opp := game[0]
		mine := game[1]

		score = score + points[mine]
		if mine == winMoves[opp] {
			score = score + winPoint
		} else if mine == drawMoves[opp] {
			score = score + drawPoint
		}
	}
	log.Printf("Score %v", score)
}
