package main

import (
	"fmt"
	"sort"
)

type mastruct struct {
	ball int
	bat  int
}

func variadicFunc(args ...int) {
	for _, v := range args {
		print(v)
	}

}
func sum(a int, b int) int { return a + b }
func fib(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}

func half(n int) (int, bool) {
	return n / 2, n%2 == 0
}

func runHalf(n int, done chan<- bool) {
	fmt.Println(half(n))
	done <- false

}
func (st mastruct) printStruct() {
	println("ball: ", st.ball)
	println("hats: ", st.bat)
}

func main() {
	done := make(chan bool)

	add := func(rhs int, lhs int) int {
		return rhs + lhs
	}

	defer variadicFunc(1, 2, 3, 4, 545)
	var lol mastruct
	lol.ball = 32
	lol.bat = 44
	arr := []int{1, 2, 3, 4}
	arr = append(arr, 22)

	add(lol.ball, lol.bat)

	fmt.Println(arr)
	fmt.Println(lol)
	lol.printStruct()

	vertices := make(map[int]string)
	vertices[3] = "sdsdsd"
	vertices[4] = "4444"
	vertices[5] = "777"

	abstractArray := []mastruct{{22, 33}, {55, 66}, {33, 22}}

	fmt.Println(abstractArray)
	sort.SliceStable(abstractArray, func(i, j int) bool {
		return abstractArray[i].ball < abstractArray[j].ball
	})
	go runHalf(56, done)

	fmt.Println(abstractArray)
	fmt.Println(vertices, add(lol.ball, lol.bat))
	fmt.Println(fib(2))

	fmt.Println(half(9))
	result := <-done

	if result {
		println("True")
	}

	fmt.Println("Second go routine is also done")

	var input string
	fmt.Scanln(&input)
}
