package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arg, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Error")
	}
	m := make(map[int]int)

	for i := 1; i <= arg; i++ {
		m[i] = i*i
	}

	fmt.Println(m)
}