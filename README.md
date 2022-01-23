# FindTheWord
This application searches given word in a given puzzle. first it will search word horizontally in puzzle and prints coordinate of first letter of word, then it will look vertically in puzzle and print  coordinate of first letter of word. After priting all occurance of word horizontally and vertically in puzzle, it will rotate puzzle by one according to rule mentioned in problem-statement.txt and again search word horizontally and vertically in rotated puzzle and print coordinate of first letter of word on first occurance.

1. [Application walkthrough](#application-walkthrough)
2. [My assumptions](#my-assumptions)
3. [Steps to run application](#steps-to-run-application)
4. [Unit test results](#unit-test-results)
5. [Improvements](#improvements)

## Application walkthrough
1. input.txt: This file contains required input (i.e. word to be searched and puzzle)
2. response.txt: This file stores result of program.
3. main.go: This file contains main() which is starting point of this application and it also includes other helper functions
4. Dockerfile: This file will be used to build image of this application.
5. problem-statement.txt: This file contains explanation of problem statement.
## My assumptions

1. Required input (i.e. word to be searched, puzzle) will be received from input.txt. First line in that file defines the word which has to be searched in puzzle.
characters in puzzle is defined from second line of input.txt. Following is the screenshot of input.txt:
![image](https://user-images.githubusercontent.com/33225945/150671478-70c9dcc2-b950-48d7-814b-162243548049.png)

2. Each row in puzzle will be of same length.
4. In problem statement it was clear whether program needs to print only first occurance of word or all occurance of word found horizontally and vertically (including rotation). So my assumptions are :
  - print <b>ALL</b> occurance of word found horizontally without rotation (add tag horizontal in result)
    output ex: abc found at x=2, y=0, count=0 - horizontal
  - print <b>ALL</b> occurance of word found vertically without rotation (add tag vertical in result)
    output ex: abc found at x=2, y=0, count=0 - vertical
  - print <b>ONLY FIRST</b> occurance of word found horizontally or vertically after rotation (add count of rotation)
    output ex: abc found at x=2, y=0, count=2 - vertical
   </br>for ex:</br>
   --------------------------- input.txt ------------------------------ </br>
   abc</br>
   a,b,c,z,a,y</br>
   x,y,z,w,b,z</br>
   p,q,r,n,c,a</br></br>
   ![image](https://user-images.githubusercontent.com/33225945/150671772-6285a47a-60ba-4226-bb5f-daabe4c13c10.png)</br></br>
   
   --------------------------- output: ---------------------------------</br>
   abc found at x=0, y=0, count=0 - horizontal</br>
   abc found at x=4, y=0, count=0 - vertical</br>
   abc found at x=5, y=0, count=2 - vertical</br>

## Steps to run application

## Unit test results

## Improvements
1. At last moment i realised all helper function defined in main.go could be moved to utils.go. main.go could have only main() defined.
2. This program can be optimised furthur by concurrently excuting some functions using goroutine. 
