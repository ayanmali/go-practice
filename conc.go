package main

import (
	"fmt"
	"math"
	"math/rand/v2"
	"sync"
)

func square(num int) int {
	return num * num
}

func cube(num int) int {
	return num * num * num
}

func root(num int) float64 {
	return math.Sqrt(float64(num))
}

func main() {
	//randomNums := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// Setting the array to 10 random integers
	var randomNums [10]int
	for i := 0; i < len(randomNums); i++ {
		randomNums[i] = rand.IntN(100)
	}

	fmt.Printf("List of Random Numbers (0-100):\n%v\n", randomNums)
	var squaredNumbers []int
	var cubedNumbers []int
	var rootNumbers []float64

	squareChan := make(chan int)
	cubeChan := make(chan int)
	rootChan := make(chan float64)

	var wg sync.WaitGroup
	//var mut sync.Mutex

	for i := 0; i < len(randomNums); i++ {
		wg.Add(3)

		go func(index int) {
			defer wg.Done()
			//var squaredNum int = square(randomNums[index])
			squareChan <- square(randomNums[index])
			//mut.Lock()
			//squaredNumbers = append(squaredNumbers, squaredNum)
			//mut.Unlock()
		}(i)

		go func(index int) {
			defer wg.Done()
			cubeChan <- cube(randomNums[index])
		}(i)

		go func(index int) {
			defer wg.Done()
			rootChan <- root(randomNums[index])
		}(i)

		squaredNum := <-squareChan
		cubedNum := <-cubeChan
		rootNum := <-rootChan

		squaredNumbers = append(squaredNumbers, squaredNum)
		cubedNumbers = append(cubedNumbers, cubedNum)
		rootNumbers = append(rootNumbers, rootNum)
	}

	wg.Wait()

	fmt.Println("Done")
	fmt.Println(squaredNumbers)
	fmt.Println(cubedNumbers)
	fmt.Println(rootNumbers)
}
