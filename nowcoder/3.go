package main

import (
	"fmt"
)

func main() {
	a := 0
	fmt.Scan(&a)
	var (
		str string
	)

	if a > 1 {
		str += fmt.Sprintf("%v ", a)
	}
	fmt.Println(str)
}
