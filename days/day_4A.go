package days

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var direction = map[string][]int{
	"East":      {0, 1},
	"SouthEast": {1, 1},
	"South":     {1, 0},
	"SouthWest": {1, -1},
	"West":      {0, -1},
	"NorthWest": {-1, -1},
	"North":     {-1, 0},
	"NorthEast": {-1, 1},
}

func RunDay4A(test bool) {
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

func findStart(board [][]string, word string) int {
	found := 0
	// Iterate through every cell in the board
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			// Start the search if the first letter matches
			if board[i][j] == string(word[0]) {
				// fmt.Printf("X Found @ %d, %d\n", i, j)
				for key, dir := range direction {
					_ = key
					isFound := searchDir(board, word, i+dir[0], j+dir[1], dir, 1)
					if isFound {
						// fmt.Printf("Found Direction: %s\n", key)
						found++
					}
				}
			}
		}
	}
	return found
}

// Recursive function to check for "XMAS"
func searchDir(board [][]string, word string, x, y int, dir []int, index int) bool {
	// If we've matched the entire word, return true

	if index == len(word) {
		return true
	}

	// Check for out-of-bounds or mismatch
	if x < 0 || y < 0 || x >= len(board) || y >= len(board[0]) || board[x][y] != string(word[index]) {
		return false
	}

	return searchDir(board, word, x+dir[0], y+dir[1], dir, index+1)
}
