package main

import (
	"fmt"
	"strconv"
	"strings"
) 

func main() {
	res := []string{}
	for i := 2000; i <= 3200; i++ {
		if i%7 == 0 && i%5 != 0 {
			res = append(res, strconv.Itoa((i)))
		}
	}
	fmt.Println(strings.Join(res, ","))
}