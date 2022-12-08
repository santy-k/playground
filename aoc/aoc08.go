package main

import (
	"log"
	"strconv"
	"strings"
)

func treeScenicScore() {
	trees, _ := initMatrices()
	maxScore := 0
	for row := 0; row < len(trees); row++ {
		for col := 0; col < len(trees[row]); col++ {
			left, right, up, down := 0, 0, 0, 0
			tree := trees[row][col]
			for j := col - 1; j >= 0; j-- {
				left++
				if trees[row][j] >= tree {
					break
				}
			}
			for j := col + 1; j < len(trees[row]); j++ {
				right++
				if trees[row][j] >= tree {
					break
				}
			}
			for i := row - 1; i >= 0; i-- {
				up++
				if trees[i][col] >= tree {
					break
				}
			}
			for i := row + 1; i < len(trees); i++ {
				down++
				if trees[i][col] >= tree {
					break
				}
			}
			score := left * right * up * down
			if score > maxScore {
				maxScore = score
			}
		}
	}

	log.Printf("maxScore: %v", maxScore)
}

func treeHouseCount() {
	trees, visibility := initMatrices()

	// For each row
	for i, row := range trees {
		tallest := -1
		for j, tree := range row { // L -> R
			if tree > tallest {
				tallest = tree
				visibility[i][j] = true
			}
		}
		tallest = -1
		for j := len(row) - 1; j >= 0; j-- { // R -> L
			if row[j] > tallest {
				tallest = row[j]
				visibility[i][j] = true
			}
		}
	}

	// For each col
	for j := 0; j < len(trees[0]); j++ {
		tallest := -1
		for i := 0; i < len(trees); i++ { // T -> B
			if trees[i][j] > tallest {
				tallest = trees[i][j]
				visibility[i][j] = true
			}
		}
		tallest = -1
		for i := len(trees) - 1; i >= 0; i-- { // B -> T
			if trees[i][j] > tallest {
				tallest = trees[i][j]
				visibility[i][j] = true
			}
		}
	}

	count := 0
	for i := 0; i < len(visibility); i++ {
		for j := 0; j < len(visibility[i]); j++ {
			if visibility[i][j] {
				count++
			}
		}
	}
	log.Printf("count %v", count)
}

func initMatrices() ([][]int, [][]bool) {
	scanner := scanner("resources/aoc_08.txt")
	var trees [][]int
	var visibility [][]bool
	for scanner.Scan() {
		heights := strings.Split(scanner.Text(), "")
		var row []int
		var visRow []bool
		for _, h := range heights {
			height, _ := strconv.Atoi(h)
			row = append(row, height)
			visRow = append(visRow, false)
		}
		visibility = append(visibility, visRow)
		trees = append(trees, row)
	}

	return trees, visibility
}
