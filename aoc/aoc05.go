package main

import (
	"fmt"
	"regexp"
	"strconv"
)

var stacks1 = [][]string{
	{"Z", "N"},
	{"M", "C", "D"},
	{"P"},
}

var stacks2 = [][]string{
	{"B", "Q", "C"},
	{"R", "Q", "W", "Z"},
	{"B", "M", "R", "L", "V"},
	{"C", "Z", "H", "V", "T", "W"},
	{"D", "Z", "H", "B", "N", "V", "G"},
	{"H", "N", "P", "C", "J", "F", "V", "Q"},
	{"D", "G", "T", "R", "W", "Z", "S"},
	{"C", "G", "M", "N", "B", "W", "Z", "P"},
	{"N", "J", "B", "M", "W", "Q", "F", "P"},
}

func crateNewTop() {
	scanner := scanner("resources/aoc_05.txt")
	stack := stacks2
	for scanner.Scan() {
		cmd := scanner.Text()
		count, from, to := parseCmd(cmd)

		var pickedItems []string
		for i := 0; i < count; i++ {
			size := len(stack[from-1])
			item := stack[from-1][size-1]
			stack[from-1] = stack[from-1][:size-1]
			pickedItems = append(pickedItems, item)
		}
		for i := len(pickedItems) - 1; i > -1; i-- {
			stack[to-1] = append(stack[to-1], pickedItems[i])
		}
	}
	for _, st := range stack {
		size := len(st)
		fmt.Print(st[size-1])
		st = st[:size-1]
	}
}

func crateTop() {
	scanner := scanner("resources/aoc_05.txt")
	stack := stacks2
	for scanner.Scan() {
		cmd := scanner.Text()
		count, from, to := parseCmd(cmd)
		for i := 0; i < count; i++ {
			size := len(stack[from-1])
			item := stack[from-1][size-1]
			stack[from-1] = stack[from-1][:size-1]
			stack[to-1] = append(stack[to-1], item)
		}
	}
	for _, st := range stack {
		size := len(st)
		fmt.Print(st[size-1])
		st = st[:size-1]
	}
}

func parseCmd(cmd string) (int, int, int) {
	re := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
	match := re.FindStringSubmatch(cmd)

	count, _ := strconv.Atoi(match[1])
	from, _ := strconv.Atoi(match[2])
	to, _ := strconv.Atoi(match[3])
	return count, from, to
}
