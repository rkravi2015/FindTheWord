package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func StoreResponse(response *[]string) {
	f, err := os.Create("response.txt")
	if err != nil {
		return
	}
	for _, res := range *response {
		_, err := f.WriteString(res + "\n")
		if err != nil {
			fmt.Println(err)
			f.Close()
			return
		}
	}
	err = f.Close()
	if err != nil {
		return
	}
}

func GetInput(inputFileName string, response *[]string) (searchWord string, puzzle [][]string) {
	fileBytes, err := ioutil.ReadFile(inputFileName)
	if err != nil {
		*response = append(*response, fmt.Sprintf("Error while reading input file. error: %s", err.Error()))
		return
	}
	fileData := strings.Split(string(fileBytes), "\n")
	searchWord = fileData[0]
	puzzle = [][]string{}
	for i := 1; i < len(fileData); i++ {
		puzzle = append(puzzle, strings.Split(fileData[i], ","))
	}
	return searchWord, puzzle
}

func SearchWordInPuzzle(word string, puzzle [][]string) []string {
	return []string{"found at 1", "found at 2"}
}

func main() {
	response := make([]string, 0)
	defer StoreResponse(&response)
	args := os.Args
	if len(args) <= 1 {
		response = append(response, "Input file name is missing.")
		return
	}
	inputFileName := args[1]
	searchWord, puzzle := GetInput(inputFileName, &response)
	searchResult := SearchWordInPuzzle(searchWord, puzzle)
	response = append(response, searchResult...)
}
