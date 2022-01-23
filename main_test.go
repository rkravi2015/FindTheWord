package main

import (
	"io/ioutil"
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestStoreResponse(t *testing.T) {
	type args struct {
		response       *[]string
		outputFileName string
	}
	tests := []struct {
		name             string
		args             args
		wordTobeSearched string
	}{
		{
			name: "success",
			args: args{
				response:       &[]string{"dummy"},
				outputFileName: "response_test.txt",
			},
			wordTobeSearched: "dummy",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			StoreResponse(tt.args.response, tt.args.outputFileName)
			fileBytes, err := ioutil.ReadFile(tt.args.outputFileName)
			if err != nil {
				t.Errorf("StoreResponse() = unable to read file: %s", tt.args.outputFileName)
			}
			fileData := strings.Split(string(fileBytes), "\n")
			if !reflect.DeepEqual(fileData[0], tt.wordTobeSearched) {
				t.Errorf("StoreResponse() word = %v, want %v", fileData[0], tt.wordTobeSearched)
			}
			_ = os.Remove(tt.args.outputFileName)
		})
	}
}

func TestGetInput(t *testing.T) {
	type args struct {
		inputFileName string
		response      *[]string
	}
	tests := []struct {
		name           string
		args           args
		wantSearchWord string
		wantPuzzle     [][]string
		responseArr    []string
	}{
		{
			name: "success",
			args: args{
				inputFileName: "input_test.txt",
				response:      &[]string{},
			},
			wantSearchWord: "abc",
			wantPuzzle: [][]string{
				{"a", "b", "c"},
				{"x", "y", "z"},
				{"p", "q", "r"},
			},
			responseArr: []string{},
		},
		{
			name: "failed - file not found",
			args: args{
				inputFileName: "dummy.txt",
				response:      &[]string{},
			},
			wantSearchWord: "",
			responseArr:    []string{"Error while reading input file. error: open dummy.txt: The system cannot find the file specified."},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSearchWord, gotPuzzle := GetInput(tt.args.inputFileName, tt.args.response)
			if gotSearchWord != tt.wantSearchWord {
				t.Errorf("GetInput() gotSearchWord = %v, want %v", gotSearchWord, tt.wantSearchWord)
			}
			if !reflect.DeepEqual(gotPuzzle, tt.wantPuzzle) {
				t.Errorf("GetInput() gotPuzzle = %v, want %v", gotPuzzle, tt.wantPuzzle)
			}
			if !reflect.DeepEqual(*tt.args.response, tt.responseArr) {
				t.Errorf("GetInput() gotResponse = %v, want %v", *tt.args.response, tt.responseArr)
			}
		})
	}
}

func TestSearchWordInPuzzle(t *testing.T) {
	type args struct {
		word   string
		puzzle [][]string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "success",
			args: args{
				word: "abc",
				puzzle: [][]string{
					{"a", "b", "c", "a"},
					{"p", "q", "r", "b"},
					{"w", "x", "y", "c"},
				},
			},
			want: []string{
				"abc found at x=0, y=0, count=0 - horizontal",
				"abc found at x=3, y=0, count=0 - vertical",
				"abc found at x=0, y=2, count=4 - horizontal",
			},
		},
		{
			name: "only horizontal search without rotation",
			args: args{
				word: "abc",
				puzzle: [][]string{
					{"a", "b", "c", "a"},
					{"x", "q", "r", "b"},
				},
			},
			want: []string{
				"abc found at x=0, y=0, count=0 - horizontal",
			},
		},
		{
			name: "only vertical search without rotation",
			args: args{
				word: "abc",
				puzzle: [][]string{
					{"a", "x", "c", "a"},
					{"x", "q", "r", "b"},
					{"q", "x", "y", "c"},
				},
			},
			want: []string{
				"abc found at x=3, y=0, count=0 - vertical",
			},
		},
		{
			name: "only horizontal search with rotation",
			args: args{
				word: "abc",
				puzzle: [][]string{
					{"a", "z", "a", "b"},
					{"x", "q", "r", "b"},
					{"q", "x", "y", "c"},
				},
			},
			want: []string{
				"abc found at x=1, y=0, count=1 - horizontal",
			},
		},
		{
			name: "only vertical search with rotation",
			args: args{
				word: "abc",
				puzzle: [][]string{
					{"a", "z", "a", "z"},
					{"x", "q", "r", "a"},
					{"q", "x", "y", "b"},
				},
			},
			want: []string{
				"abc found at x=3, y=0, count=1 - vertical",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SearchWordInPuzzle(tt.args.word, tt.args.puzzle); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SearchWordInPuzzle() = %v, want %v", got, tt.want)
			}
		})
	}
}
