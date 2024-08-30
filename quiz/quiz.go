package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
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

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}
	fmt.Println("CSV file loaded")

	// Creating the map to store questions and their answers
	var questions map[string]int = make(map[string]int)

	// Storing each question and its corresponding answer in the map
	for _, record := range records {
		intVal, err := strconv.Atoi(record[1])
		if err == nil {
			questions[record[0]] = intVal
		}
	}

	//fmt.Println(questions)
}
