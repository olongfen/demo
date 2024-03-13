package main

import "fmt"

func main() {
	fmt.Println(reverse(123))
}

func reverse(x int) int {
	num := 0
	for x != 0 {
		n := x % 10
		num = num*10 + n
		maxInt := 1<<31 - 1
		minInt := -1 << 31
		x = x / 10
		if num > maxInt || num < minInt {
			return 0
		}

	}
	return num
}
