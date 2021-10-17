package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
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

	results := []chan int{}

	for count < len(intInput) {
		result := make(chan int)
		results = append(results, result)

		subSlice := intInput[prevCount:count]
		fmt.Println(subSlice)
		go getOdd(subSlice, result)

		prevCount = count
		count += 2
	}

	for _, result := range results {
		out := <-result
		tot += out
	}

	var lastIndex int = len(intInput)
	if len(intInput)%2 != 0 && intInput[lastIndex-1]%2 != 0 {
		tot++
	}

	fmt.Println(tot)
	var input string
	fmt.Scanln(&input)

}
