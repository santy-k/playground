package main

import (
	"bufio"
	"log"
	"sort"
	"strconv"
	"strings"
)

var total = 0
var dirSizes []int

func smallestDirToDelete() {
	totalSize := totalSize()
	fileSizes(scanner("resources/aoc_07.txt"))
	freeSpace := 70000000 - totalSize
	spaceToClear := 30000000 - freeSpace

	sort.Ints(dirSizes)
	for _, space := range dirSizes {
		if space >= spaceToClear {
			log.Printf("smallestDirToDelete: %v", space)
			break
		}
	}
}

func totalSizeOfSmallDirs() {
	fileSizes(scanner("resources/aoc_07.txt"))
	log.Printf("total: %v", total)
}

func fileSizes(scanner *bufio.Scanner) int {
	size := 0
	for scanner.Scan() {
		text := scanner.Text()
		if _, err := strconv.Atoi(text[0:1]); err == nil { // is a file
			fileSize, _ := strconv.Atoi(strings.Split(text, " ")[0])
			size = size + fileSize
		} else if text[0:4] == "dir " {
			// is a dir, do nothing
		} else if text[0:4] == "$ cd" { // change dir
			dir := strings.Split(text, " ")[2] // dir name
			if dir == ".." {
				if size <= 100000 {
					total = total + size
				}
				dirSizes = append(dirSizes, size)
				return size
			} else {
				size = size + fileSizes(scanner)
			}
		}
	}
	if size <= 100000 {
		total = total + size
	}
	dirSizes = append(dirSizes, size)
	return size
}

func totalSize() int {
	scanner := scanner("resources/aoc_07.txt")
	size := 0
	for scanner.Scan() {
		text := scanner.Text()
		if _, err := strconv.Atoi(text[0:1]); err == nil { // is a file
			fileSize, _ := strconv.Atoi(strings.Split(text, " ")[0])
			size = size + fileSize
		}
	}
	return size
}
