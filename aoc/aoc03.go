package main

import (
	"log"
)

func elfGroup() {
	scanner := scanner("resources/aoc_03.txt")
	sum := 0
	count := 0
	var items = [3]string{}
	for scanner.Scan() {
		count = count + 1
		items[count-1] = scanner.Text()
		if count%3 == 0 {
			elf1Set := make(map[int]bool)
			elf2Set := make(map[int]bool)
			elf3Set := make(map[int]bool)

			for _, c := range items[0] {
				elf1Set[int(c)] = true
			}
			for _, c := range items[1] {
				elf2Set[int(c)] = true
			}
			for _, c := range items[2] {
				elf3Set[int(c)] = true
			}
			for c := range elf3Set {
				if elf1Set[c] && elf2Set[c] {
					if c >= 97 {
						sum = sum + c - 96
					} else {
						sum = 26 + sum + c - 64
					}
				}
			}
			count = 0
		}
	}
	log.Printf("Sum: %v", sum)
}

func elfPackage() {
	scanner := scanner("resources/aoc_03.txt")
	sum := 0

	for scanner.Scan() {
		items := scanner.Text()
		length := len(items)
		cmp1Set := make(map[int]bool)
		cmp2Set := make(map[int]bool)

		for _, c := range items[0 : length/2] {
			cmp1Set[int(c)] = true
		}
		for _, c := range items[length/2:] {
			cmp2Set[int(c)] = true
		}
		for c := range cmp1Set {
			if cmp2Set[c] {
				if c >= 97 {
					sum = sum + c - 96
				} else {
					sum = 26 + sum + c - 64
				}
			}
		}
	}
	log.Printf("Sum: %v", sum)
}
