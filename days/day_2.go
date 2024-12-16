package days

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
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
	safeReports := 0
	reportNum := 0
	for scanner.Scan() {
		reportStr := strings.Split(scanner.Text(), " ")
		safe, _ := processReport(reportStr)
		if safe {
			safeReports++
		} else {
			for i := 0; i < len(reportStr); i++ {
				cprep := append([]string{}, reportStr...)
				report := slices.Delete(cprep, i, i+1)
				safe, _ := processReport(report)
				fmt.Println(reportStr, cprep, report, safe)
				if safe {
					safeReports++
					break
				}
			}
		}
		reportNum++
	}
	fmt.Printf("Safe Reports: %d \n", safeReports)
	return
}

func processReport(str []string) (bool, int) {
	report := []int{}
	for _, val := range str {
		intVal, _ := strconv.Atoi(val)
		report = append(report, intVal)
	}
	// safe := true
	direction := 0
	if report[0] > report[1] {
		direction = -1
	} else if report[0] < report[1] {
		direction = 1
	} else {
		return false, 0
	}

	for i := 1; i < len(report); i++ {
		// fmt.Printf("Direction: %v \n", direction)
		diff := math.Abs(float64(report[i-1]) - float64(report[i]))
		// fmt.Printf("Diff: %v \n", diff)
		if diff <= 3 && diff >= 1 {
			switch direction {
			case -1:
				if report[i-1] < report[i] {
					// fmt.Printf("Fail: %v > %v Round: %v \n", report[i-1], report[i], i)
					return false, i
				}
			case 1:
				if report[i-1] > report[i] {
					// fmt.Printf("Fail: %v < %v Round: %v \n", report[i-1], report[i], i)
					return false, i
				}
			default:
				// fmt.Printf("Fail: Huh?  %v %v Round: %v \n", report[i-1], report[i], i)
				return false, i
			}
		} else {
			// fmt.Printf("Fail: %v %v Round: %v \n", report[i-1], report[i], i)
			return false, i
		}
	}
	return true, 0
}
