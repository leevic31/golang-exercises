package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println(factorial(8))
	fmt.Println(factorial2())
}

func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return factorial(n-1) * n
}

func factorial2() int {
	arg, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Error")
	}
	res := 1
	for i := arg; i > 0; i-- {
		res *= i
	}
	return res
}