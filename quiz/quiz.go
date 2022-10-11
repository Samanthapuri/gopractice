package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"strings"
)

var userinput string
var answer string

func main() {
	data, err := fs.ReadFile(os.DirFS("/Users/vijaysamanthapuri/Documents/gopractice/quiz"), "problems.csv")
	if err != nil {
		log.Fatal("Error while reading file")
	}
	fmt.Println("\nPress ENNTER to start quiz\n")
	fmt.Scanf("%s", &userinput)
	//fmt.Println(userinput)
	if userinput != "" {
		log.Fatal("Invalid input!!")
	}

	score := 0

	for _, val := range strings.Split(string(data), "\n") {

		fmt.Print(strings.Split(val, ",")[0] + " : ")
		fmt.Scanf("%s", &answer)
		if realanswer := strings.Split(val, ",")[1]; answer == realanswer {
			score++
		}
	}
	fmt.Printf("\n\n You have completed the exam. Your score is %v\n\n", score)
}
