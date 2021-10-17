package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"sync"
)

func getOdd(vect []int, result chan<- int) {
	c := 0
	for _, v := range vect {
		if v%2 != 0 {
			c++
		}

	}
	result <- c
}

func main() {
	inputData, err := ioutil.ReadFile("data.txt")

	if err != nil {
		fmt.Println("error while opening file: ", nil)
	}

	sliceData := strings.Split(string(inputData), ",")

	intInput := []int{}

	for _, v := range sliceData {
		intVal, err := strconv.Atoi(v)

		if err != nil {
			fmt.Println("error while parsing data: ", nil)
		}
		intInput = append(intInput, intVal)
	}

	var count int = 2
	var prevCount int = 0
	var tot int = 0

	results := make(chan int, len(intInput))
	var wg sync.WaitGroup

	for count < len(intInput) {
		subSlice := intInput[prevCount:count]
		fmt.Println(subSlice)
		wg.Add(1)
		go getOdd(subSlice, results)

		prevCount = count
		count += 2
	}

	go func() {
		for result := range results {
			tot += result
			wg.Done()
		}
	}()

	wg.Wait()
	fmt.Println(tot)

}
