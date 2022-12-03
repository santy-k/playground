package main

import (
	"log"
	"sort"
	"strconv"
)

func elfCaloriesTop3Sum() {
	scanner := scanner("resources/aoc_01.txt")
	var totalCalories []int
	cur := 0
	for scanner.Scan() {
		text := scanner.Text()
		if text != "" {
			calorie, _ := strconv.Atoi(text)
			cur = cur + calorie
		} else {
			totalCalories = append(totalCalories, cur)
			cur = 0
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(totalCalories)))
	sum := 0
	for _, n := range totalCalories[0:3] {
		sum = sum + n
	}

	log.Printf("Top 3 sum: %v", sum)
}

func elfMostCalories() {
	scanner := scanner("resources/aoc_01.txt")
	max := 0
	cur := 0
	for scanner.Scan() {
		text := scanner.Text()
		if text != "" {
			calorie, _ := strconv.Atoi(text)
			cur = cur + calorie
			if cur > max {
				max = cur
			}
		} else {
			cur = 0
		}
	}

	log.Printf("Max: %v", max)
}
