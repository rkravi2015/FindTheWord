package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const (
	//outputFileName is name of file in which response will be stored
	outputFileName = "response.txt"
	zChar = 'z'
	aChar = 'a'
)

func main() {
	// response stores error messages and search result which will be shown to user after successful run of this program
	response := make([]string, 0)

	// defer will call StoreResponse to store output captured from this program in a .txt file
	defer StoreResponse(&response, outputFileName)

	// read command line argument
	args := os.Args
	if len(args) <= 1 {
		response = append(response, "Input file name is missing.")
		return
	}

	//read input filename from command line argument array
	inputFileName := args[1]

	//fetch puzzle and word to be searched from input file
	searchWord, puzzle := GetInput(inputFileName, &response)

	//search given in puzzle
	searchResult := SearchWordInPuzzle(searchWord, puzzle)

	//after searching store result in resposne array
	response = append(response, searchResult...)
}

//StoreResponse stores list of string in a .txt file
func StoreResponse(response *[]string, outputFileName string) {
	//create .txt file to store response
	f, err := os.Create(outputFileName)
	if err != nil {
		return
	}

	//write item given in response array to output file
	for _, res := range *response {
		_, err := f.WriteString(res + "\n")
		if err != nil {
			fmt.Println(err)
			f.Close()
			return
		}
	}

	// close file
	err = f.Close()
	if err != nil {
		return
	}
}

//GetInput read input file and returns puzzle and word to be searched
func GetInput(inputFileName string, response *[]string) (searchWord string, puzzle [][]string) {

	//read input file name
	fileBytes, err := ioutil.ReadFile(inputFileName)
	if err != nil {
		*response = append(*response, fmt.Sprintf("Error while reading input file. error: %s", err.Error()))
		return
	}
	fileData := strings.Split(string(fileBytes), "\n")

	//first line in input file is word to be searched
	searchWord = strings.Trim(fileData[0], " \r\n")

	//from second line and onwards in input file defines puzzle
	puzzle = [][]string{}
	for i := 1; i < len(fileData); i++ {
		puzzle = append(puzzle, strings.Split(strings.Trim(fileData[i], " \r\n"), ","))
	}
	return searchWord, puzzle
}

//SearchWordInPuzzle searches given word in given puzzle
func SearchWordInPuzzle(word string, puzzle [][]string) []string {

	//response stores search result
	response := make([]string, 0)

	//search row wise without rotation
	rowWiseResponse := SearchWordRowWise(word, puzzle, 0)
	response = append(response, rowWiseResponse...)

	//search column wise without rotation
	colWiseResponse := SearchWordColumnWise(word, puzzle, 0)
	response = append(response, colWiseResponse...)

	//count stores rotation count
	count := 0

	//rotate puzzle and search
	for count < 25 {
		puzzle = RotatePuzzle(puzzle)
		count = count + 1
		rowWiseResponse := SearchWordRowWise(word, puzzle, count)
		if len(rowWiseResponse) > 0 {
			response = append(response, rowWiseResponse...)
			break
		}
		colWiseResponse := SearchWordColumnWise(word, puzzle, count)
		if len(colWiseResponse) > 0 {
			response = append(response, colWiseResponse...)
			break
		}
	}
	return response
}

//SearchWordRowWise searches given word in every row of given puzzle
func SearchWordRowWise(searchWord string, puzzle [][]string, rotationCount int) []string {
	response := make([]string, 0)
	for yIndex, value := range puzzle {
		rowWord := strings.Join(value, "")
		if strings.Contains(rowWord, searchWord) {
			xIndex := strings.Index(rowWord, searchWord)
			response = append(response, fmt.Sprintf("%s found at x=%d, y=%d, count=%d - horizontal", searchWord, xIndex, yIndex, rotationCount))
		}
	}
	return response
}

//SearchWordColumnWise searches given word every column of given puzzle
func SearchWordColumnWise(searchWord string, puzzle [][]string, rotationCount int) []string {
	response := make([]string, 0)
	for col := 0; col < len(puzzle[0]); col++ {
		colWord := ""
		for row := 0; row < len(puzzle); row++ {
			colWord = fmt.Sprintf("%s%s", colWord, puzzle[row][col])
		}
		if strings.Contains(colWord, searchWord) {
			yIndex := strings.Index(colWord, searchWord)
			response = append(response, fmt.Sprintf("%s found at x=%d, y=%d, count=%d - vertical", searchWord, col, yIndex, rotationCount))
		}
	}
	return response
}

//Increment returns next character of given input
func Increment(item string) string {
	if item == "" || len(item) > 1 {
		return item
	}
	itemRune := []rune(item)
	charToIncrement := itemRune[0]
	if charToIncrement < zChar {
		return string(charToIncrement + 1)
	}
	return string(aChar)
}

//RotatePuzzle returns rotated puzzle
func RotatePuzzle(puzzle [][]string) [][]string {
	for rowIndex, row := range puzzle {
		for colIndex, _ := range row {
			puzzle[rowIndex][colIndex] = Increment(puzzle[rowIndex][colIndex])
		}
	}
	return puzzle
}
