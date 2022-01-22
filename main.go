package main

import (
	"FindTheWord/utils"
	"os"
)

func main() {
	response := make([]string, 0)
	defer utils.StoreResponse(&response)
	args := os.Args
	if len(args) <= 1 {
		response = append(response, "Input file name is missing.")
		return
	}
	inputFileName := args[1]
	searchWord, puzzle := utils.GetInput(inputFileName, &response)
	//fmt.Println("***********")
	//for _, v:= range puzzle {
	//	fmt.Println(v)
	//}
	//fmt.Println("***********")
	searchResult := utils.SearchWordInPuzzle(searchWord, puzzle)
	response = append(response, searchResult...)
}
