package main

import (
	days "aoc/days"
	"aoc/utils"
	_ "aoc/utils"
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(".env file not found.")
	}
	isTest := false
	if len(os.Args) > 1 && os.Args[1] == "test" {
		isTest = true
	}
	days := map[string]func(test bool){
		"1": days.RunDay1,
		"2": days.RunDay2,
		"3": days.RunDay3,
		"4": days.RunDay4A,
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Advent Of Code")

	for i := range days {
		fmt.Printf("Day %s\n", i)
	}

	fmt.Println("C: Create New Day")
	fmt.Println("Q: Quit")
	choice := utils.PromptUser("Enter day number: ", reader)
	choice = strings.TrimSpace(choice)
	choice = strings.ToLower(choice)

	if choice == "q" {
		fmt.Println("Exiting...")
		os.Exit(0)
	} else if choice == "c" {

		newDay := utils.PromptUser("Enter day number: ", reader)
		utils.CreateDay(newDay)
	} else {
		days[choice](isTest)
	}
}
