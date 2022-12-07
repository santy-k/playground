package main

import (
	"log"
	"strconv"
	"strings"
)

type interval struct {
	start int
	end   int
}

func overLappingdWork() {
	scanner := scanner("resources/aoc_04.txt")
	count := 0
	for scanner.Scan() {
		rangesStr := strings.Split(scanner.Text(), ",")
		interval1 := createInterval(rangesStr[0])
		interval2 := createInterval(rangesStr[1])

		if interval1.start > interval2.start { // interval1 should be the first one
			interval1, interval2 = interval2, interval1
		}

		if interval1.end >= interval2.start {
			count = count + 1
		}
	}
	log.Printf("count: %v", count)
}

func embeddedWork() {
	scanner := scanner("resources/aoc_04.txt")
	count := 0
	for scanner.Scan() {
		rangesStr := strings.Split(scanner.Text(), ",")
		interval1 := createInterval(rangesStr[0])
		interval2 := createInterval(rangesStr[1])

		if interval1.end-interval1.start < interval2.end-interval2.start { // interval1 should be bigger
			interval1, interval2 = interval2, interval1
		}

		if interval1.start <= interval2.start && interval1.end >= interval2.end {
			count = count + 1
		}
	}
	log.Printf("count: %v", count)
}

func createInterval(str string) interval {
	intervalStr := strings.Split(str, "-")
	start, _ := strconv.Atoi(intervalStr[0])
	end, _ := strconv.Atoi(intervalStr[1])
	return interval{start: start, end: end}
}
