package main

import (
	"log"
	"strings"
)

func startOfPacket(uniqCount int) {
	scanner := scanner("resources/aoc_06.txt")
	for scanner.Scan() {
		input := scanner.Text()
		chars := strings.Split(input, "")
		uniqChars := map[string]bool{}
		charsCount := map[string]int{}

		for i := 0; i < uniqCount-1; i++ {
			char := chars[i]
			uniqChars[char] = true
			addChar(charsCount, char)
		}
		i := 0
		for j := uniqCount - 1; j < len(chars); j++ {
			begin := chars[i]
			end := chars[j]

			// add end char
			uniqChars[end] = true
			addChar(charsCount, end)

			// check if found
			if len(uniqChars) == uniqCount {
				log.Printf("Found: %v", j+1)
				break
			}

			// remove begin character
			val, _ := charsCount[begin]
			if val == 1 {
				delete(charsCount, begin)
				delete(uniqChars, begin)
			} else {
				charsCount[begin] = val - 1
			}
			i = i + 1
		}
	}
}

func addChar(charMap map[string]int, char string) {
	val, ok := charMap[char]
	if ok {
		charMap[char] = val + 1
	} else {
		charMap[char] = 1
	}
}
