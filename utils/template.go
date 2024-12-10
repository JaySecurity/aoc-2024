package utils

import (
	"fmt"
	"os"
	"regexp"
)

func CreateDay(day string) {
	template := `
package days

import (
  "aoc/utils"
  "os"
  "fmt"
)

func RunDay%day%(test bool){

  // utils.GetData("%day%")
	var file string

	if test {
		file = "./day_%day%/sample.txt"
	} else {
		file = "./day_%day%/input.txt"
	}

	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Printf("Error Reading FIle: %v", err)
	}
  return 
}

`

	pattern := `%day%`
	templateBytes := []byte(template)
	re := regexp.MustCompile(pattern)
	data := re.ReplaceAll(templateBytes, []byte(day))
	filePath := fmt.Sprintf("./day_%s", day)
	fileName := fmt.Sprintf("./days/day_%s.go", day)
	os.Mkdir(filePath, 0755)
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Write the response body to the file
	_, err = file.Write(data)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	GetData(day)
}
