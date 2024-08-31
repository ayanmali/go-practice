package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	const filename string = "problems.csv"

	// Opening the file
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	fmt.Printf("File %v opened successfully", filename)
	// Closing the file once the code is finished executing
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		panic(err)
	}
	fmt.Println("CSV file loaded")

	// Creating the map to store questions and their answers
	var questions map[string]string = make(map[string]string)

	// Storing each question and its corresponding answer in the map
	for _, record := range records {
		answer := record[1]
		questions[record[0]] = answer
	}

	// fmt.Println(questions)

	// Creating the reader object
	inputReader := bufio.NewReader(os.Stdin)
	// initializing a counter variable to keep track of score
	var correctAnswers int = 0
	var total int = len(questions)
	answerCh := make(chan string)
	// Creating the timer
	timer := time.NewTimer(30 * time.Second)

	// Display each question and get user input for the answer
	for key := range questions {
		// Displays the given question
		fmt.Printf("%s\n", key)
		go func() {
			// Gets the users answer
			userAnswer, _ := inputReader.ReadString('\n')
			// Removing the newline character from the user answer string
			//userAnswer = userAnswer[:len(userAnswer)-1]
			userAnswer = strings.TrimSpace(userAnswer)

			fmt.Printf("You answered %s\n", userAnswer)
			fmt.Printf("The correct answer is %s\n", questions[key])

			answerCh <- userAnswer
		}()

		select {
		case <-timer.C:
			fmt.Println("Time is up")
			fmt.Printf("Your final score is: %v / %v", correctAnswers, total)
			return
		case userAnswer := <-answerCh:
			// Tells the user if they got it right or not
			if userAnswer == questions[key] {
				fmt.Println("Correct")
				correctAnswers++
			} else {
				fmt.Println("Incorrect")
			}
		}

		fmt.Println()
	}

	fmt.Printf("Your final score is: %v / %v", correctAnswers, total)

}
