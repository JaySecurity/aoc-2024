package days

import (
	"bufio"
	"fmt"
	"os"
)

func RunDay2(test bool) {
	// utils.GetData("2")
	var file string

	if test {
		file = "./day_2/sample.txt"
	} else {
		file = "./day_2/input.txt"
	}

	f, err := os.Open(file)
	if err != nil {
		fmt.Printf("Error Reading FIle: %v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		// do something with a line
		fmt.Printf("line: %s\n", scanner.Text())
	}
	return
}
