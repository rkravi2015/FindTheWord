# FindTheWord
This application searches given word in a given puzzle. first it will search word horizontally in puzzle and prints coordinate of first letter of word, then it will look vertically in puzzle and print  coordinate of first letter of word. After priting all occurance of word horizontally and vertically in puzzle, it will rotate puzzle by one according to rule mentioned in instruction and again search word horizontally and vertically in rotated puzzle and print coordinate of first letter of word on first occurance.

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
5. main_test.go: Test file for main.go.
6. input_test.txt: Test file used for unit testing.

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
   ```
   --------------------------- input.txt ------------------------------
   abc
   a,b,c,z,a,y
   x,y,z,w,b,z
   p,q,r,n,c,a
   ```
   
   ![image](https://user-images.githubusercontent.com/33225945/150671772-6285a47a-60ba-4226-bb5f-daabe4c13c10.png)</br></br>
   ```
   --------------------------- output: ---------------------------------</br>
   abc found at x=0, y=0, count=0 - horizontal</br>
   abc found at x=4, y=0, count=0 - vertical</br>
   abc found at x=5, y=0, count=2 - vertical</br>
   ```

## Steps to run application
This application can run on any system having docker installed.
My recommendation is to use "docker playground" (https://labs.play-with-docker.com/), this is a free platform to run docker container.
</br><b>note:</b> i have already pushed a docker image to docker hub with sample <a href="https://github.com/rkravi2015/FindTheWord/blob/main/input.txt" target="_blank">input.txt</a> file.
</br>image name: <b>rkravi1502/find-the-word</b>

<b>Follow following steps to run application on sample input given in <a href="https://github.com/rkravi2015/FindTheWord/blob/main/input.txt" target="_blank">input.txt</a>:</b>
 1. login to https://labs.play-with-docker.com/
 2. create an instance
 3. Pull "rkravi1502/find-the-word" image</br>
      </br>command: ```docker pull rkravi1502/find-the-word```</br>
      
      ![image](https://user-images.githubusercontent.com/33225945/150676837-43094939-2344-47cf-8f2d-2d52f4f3988d.png)

 4. Run container</br>
      </br>command: ```docker run --name c1 find-the-word``` </br>
      
      ![image](https://user-images.githubusercontent.com/33225945/150676851-6a50308a-f4be-4ecf-b468-56cfe9b057ab.png)
      
<b>Follow following steps to run application with different input file:</b>
  1. login to https://labs.play-with-docker.com/
  2. create an instance
  3. Clone repository</br>
      </br>command: ```git clone https://github.com/rkravi2015/FindTheWord.git```</br>
      
      ![image](https://user-images.githubusercontent.com/33225945/150677056-f6b6623a-dffe-43c9-bb5c-a7b3592ae938.png)

   4. Navigate to /FIndTheWord folder</br>
      </br>command: ```cd FindTheWord/```</br>
      
   5. Edit input.txt. Give your own inputs (i.e. word to be searched, puzzle ). 
      - Keep word to be searched in first row.
      - Give input for puzzle from second row.
        
        ![image](https://user-images.githubusercontent.com/33225945/150678227-f58872d8-1dd1-43e3-be16-c14b44e523e6.png)

      
   6. Build image</br>
      </br>command: ```docker build –tag find-the-word .```</br>
      
      </br><b>Note:</b> if you are using “play with docker” you might get following error:
      
      ![image](https://user-images.githubusercontent.com/33225945/150677157-ad84a0a5-1345-4768-a141-6115b9d88d4c.png)
      
      Resolution: login to docker
      </br>command: ```docker login –username <your username>```</br>
      
      Enter your password
      
      ![image](https://user-images.githubusercontent.com/33225945/150677215-b55f3d01-82cd-4fc2-9304-eff44b0d1802.png)
      
      For more details on docker login, checkout this link: https://docs.docker.com/engine/reference/commandline/login/
      If you are not an existing user then create an account using this link: https://hub.docker.com/
      
   7. Run container</br>
      </br>command: ```docker run –name c1 find-the-word```</br>
      
      ![image](https://user-images.githubusercontent.com/33225945/150677305-6be6a6ac-0c23-4ebf-bc19-fa0c67beb80f.png)


   
## Unit test results

<b>Output:</b></br>
  ```PS C:\Users\rkumar3\go\src\FindTheWord> go test -v
=== RUN   TestStoreResponse
=== RUN   TestStoreResponse/success
--- PASS: TestStoreResponse (0.00s)
    --- PASS: TestStoreResponse/success (0.00s)
=== RUN   TestGetInput
=== RUN   TestGetInput/success
=== RUN   TestGetInput/failed_-_file_not_found
--- PASS: TestGetInput (0.00s)
    --- PASS: TestGetInput/success (0.00s)
    --- PASS: TestGetInput/failed_-_file_not_found (0.00s)
=== RUN   TestSearchWordInPuzzle
=== RUN   TestSearchWordInPuzzle/success
=== RUN   TestSearchWordInPuzzle/only_horizontal_search_without_rotation
=== RUN   TestSearchWordInPuzzle/only_vertical_search_without_rotation
=== RUN   TestSearchWordInPuzzle/only_horizontal_search_with_rotation
=== RUN   TestSearchWordInPuzzle/only_vertical_search_with_rotation
--- PASS: TestSearchWordInPuzzle (0.00s)
    --- PASS: TestSearchWordInPuzzle/success (0.00s)
    --- PASS: TestSearchWordInPuzzle/only_horizontal_search_without_rotation (0.00s)
    --- PASS: TestSearchWordInPuzzle/only_vertical_search_without_rotation (0.00s)
    --- PASS: TestSearchWordInPuzzle/only_horizontal_search_with_rotation (0.00s)
    --- PASS: TestSearchWordInPuzzle/only_vertical_search_with_rotation (0.00s)
PASS
ok      FindTheWord     0.603s
```
<b> Test Cases: </b></br>

<b>GetInput(inputFileName string, response *[]string) (searchWord string, puzzle [][]string)</b>

![image](https://user-images.githubusercontent.com/33225945/150679208-f68cc070-bcf4-4525-862e-c116ee9f8794.png)

<b>SearchWordInPuzzle(word string, puzzle [][]string) []string</b>

![image](https://user-images.githubusercontent.com/33225945/150679598-2f350cd2-e6de-43e0-820a-ee5d3168bea8.png)

<b>StoreResponse(response *[]string, outputFileName string)</b>

![image](https://user-images.githubusercontent.com/33225945/150679995-1a17859e-4e6f-45eb-a274-5f741c7df84a.png)

## Improvements
1. All helper function defined in main.go could be moved to utils.go. main.go could have only main() defined.
2. This program can be optimised furthur by concurrently executing some functions using goroutine. 
