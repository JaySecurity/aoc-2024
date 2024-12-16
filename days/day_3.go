package days

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func RunDay3(test bool) {

	// utils.GetData("3")
	var file string

	if test {
		file = "./day_3/sample.txt"
	} else {
		file = "./day_3/input.txt"
	}

	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Printf("Error Reading FIle: %v", err)
	}
	pattern := `mul\((\d{1,3}),(\d{1,3})\)`

	re := regexp.MustCompile(pattern)
	matches := re.FindAllStringSubmatch(string(data), -1)
	fmt.Println(matches)
	if matches == nil {
		fmt.Println("No Matching Data Found")
		return
	}
	total := 0
	for _, match := range matches {
		fmt.Printf("Found: %v, D1: %v, D2: %v\n", match[0], match[1], match[2])
		num1, _ := strconv.Atoi(match[1])
		num2, _ := strconv.Atoi(match[2])
		total += num1 * num2
	}

	fmt.Printf("Total: %v", total)
	return
}
