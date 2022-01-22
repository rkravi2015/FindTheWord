package utils

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
	searchWord = strings.Trim(fileData[0], " \r\n")
	puzzle = [][]string{}
	for i := 1; i < len(fileData); i++ {
		puzzle = append(puzzle, strings.Split(strings.Trim(fileData[i], " \r\n"), ","))
	}
	return searchWord, puzzle
}

func SearchWordInPuzzle(word string, puzzle [][]string) []string {
	response:=make([]string, 0)
	rowWiseResponse:= searchWordRowWise(word, puzzle)
	response = append(response, rowWiseResponse...)
	colWiseResponse:= searchWordColumnWise(word, puzzle)
	response = append(response, colWiseResponse...)

	return response
}


func searchWordRowWise(searchWord string, puzzle [][]string) []string {
	response:=make([]string, 0)
	for yIndex, value:= range puzzle {
		rowWord:=strings.Join(value, "")
		if strings.Contains(rowWord, searchWord) {
			xIndex:=strings.Index(rowWord, searchWord)
			response = append(response, fmt.Sprintf("%s found at x=%d, y=%d, count=0 - horizontal", searchWord, xIndex, yIndex))
		}
	}
	return response
}

func searchWordColumnWise(searchWord string, puzzle [][]string) []string {
	response:=make([]string, 0)
	for col:=0; col<len(puzzle[0]); col++ {
		colWord:=""
		for row:=0; row<len(puzzle); row++ {
			colWord = fmt.Sprintf("%s%s", colWord, puzzle[row][col])
		}
		if strings.Contains(colWord, searchWord) {
			yIndex:=strings.Index(colWord, searchWord)
			response = append(response, fmt.Sprintf("%s found at x=%d, y=%d, count=0 - vertical", searchWord, col, yIndex))
		}
	}
	return response
}

func Increment(r rune) rune {
	if r < 'z' {
		return r + 1
	}
	return 'a'
}
