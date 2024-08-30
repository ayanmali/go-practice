package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
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

	// Display each question and get user input for the answer
	for key := range questions {
		// Displays the given question
		fmt.Printf("%s\n", key)
		// Gets the users answer
		userAnswer, _ := inputReader.ReadString('\n')
		// Removing the newline character from the user answer string
		//userAnswer = userAnswer[:len(userAnswer)-1]
		userAnswer = strings.TrimSpace(userAnswer)

		fmt.Printf("You answered %s\n", userAnswer)
		fmt.Printf("The correct answer is %s\n", questions[key])

		// Tells the user if they got it right or not
		if userAnswer == questions[key] {
			fmt.Println("Correct")
		} else {
			fmt.Println("Incorrect")
		}
		fmt.Println()
	}

}
