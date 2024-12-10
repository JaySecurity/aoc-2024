package days

import (
	"aoc/utils"
	"fmt"
	"math"
	"os"
	"regexp"
	"slices"
	"strconv"
)

func RunDay1(test bool) {
	testResult := 11
	var file string

	if test {
		file = "./day_1/sample.txt"
	} else {
		file = "./day_1/input.txt"
	}

	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Printf("Error Reading FIle: %v", err)
	}

	pattern := `(\d+)\s+(\d+)`

	re := regexp.MustCompile(pattern)
	matches := re.FindAllStringSubmatch(string(data), -1)
	if matches == nil {
		fmt.Println("No Matching Data Found")
		return
	}
	var list1, list2 []int
	for _, match := range matches {
		// Convert the matched groups to integers
		if len(match) >= 3 { // Ensure there are at least two groups
			num1, _ := strconv.Atoi(match[1])
			num2, _ := strconv.Atoi(match[2])
			list1 = append(list1, num1)
			list2 = append(list2, num2)
		}
	}

	slices.Sort(list1)
	slices.Sort(list2)
	total := float64(0)

	for i, num := range list1 {
		diff := math.Abs(float64(num) - float64(list2[i]))
		total += diff
	}

	utils.GetData("1")

	if test {
		if testResult == int(total) {
			fmt.Printf("Result: %v, PASS\n", total)
		} else {
			fmt.Printf("Result: %v, FAIL\n", total)
		}
	}
	fmt.Printf("Result: %d\n", int(total))

	part2(list1, list2, test)

	return
}

func part2(list1, list2 []int, test bool) {
	testResult := 31

	occurances := make(map[string]int)
	total := 0
	for _, item := range list1 {
		if val, ok := occurances[strconv.Itoa(item)]; ok {
			total += item * val
		} else {
			occurances[strconv.Itoa(item)] = 0
			for _, val := range list2 {
				if val == item {
					occurances[strconv.Itoa(item)]++
				}
			}
			total += item * occurances[strconv.Itoa(item)]
		}
	}

	if test {
		if testResult == int(total) {
			fmt.Printf("Result: %v, PASS\n", total)
		} else {
			fmt.Printf("Result: %v, FAIL\n", total)
		}
	}
	fmt.Printf("Result: %d\n", int(total))
}
