package days

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func RunDay4B(test bool) {
	// utils.GetData("4")
	var file string

	if test {
		file = "./day_4/sample.txt"
	} else {
		file = "./day_4/input.txt"
	}

	f, err := os.Open(file)
	if err != nil {
		fmt.Printf("Error Reading FIle: %v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var board [][]string
	for scanner.Scan() {
		line := scanner.Text()
		letters := strings.Split(line, "")
		board = append(board, letters)
	}
	found := findStart(board, "XMAS")

	fmt.Println("Found: ", found)

	return
}
