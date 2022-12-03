package main

import (
	"bufio"
	"os"
)

func scanner(filename string) *bufio.Scanner {
	file, _ := os.Open(filename)
	return bufio.NewScanner(file)
}
