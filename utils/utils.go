package utils

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func GetData(day string) {
	URL := fmt.Sprintf("https://adventofcode.com/2024/day/%s/input", day)
	client := &http.Client{}
	req, _ := http.NewRequest("GET", URL, nil)
	sessionId := os.Getenv("SESSION")
	req.Header.Set("COOKIE", fmt.Sprintf("session=%s", sessionId))
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	// Open a file to save the response body
	path := fmt.Sprintf("./day_%s/input.txt", day)
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Write the response body to the file
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}

func PromptUser(prompt string, r *bufio.Reader) string {
	fmt.Print(prompt)
	input, _ := r.ReadString('\n')
	input = strings.TrimSpace(input)
	return input
}

func ClearScreen() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default: // Assume Unix-like
		fmt.Print("\033[2J\033[H")
	}
}
