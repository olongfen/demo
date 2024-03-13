package main

import (
	"fmt"
)

func main() {
	fmt.Println(romanToInt("LVIII"))
}

func romanToInt(s string) int {
	m := map[string]int{
		"I":  1,
		"IV": 4,
		"V":  5,
		"IX": 9,
		"X":  10,
		"XL": 40,
		"L":  50,
		"XC": 90,
		"C":  100,
		"CD": 400,
		"D":  500,
		"CM": 900,
		"M":  1000,
	}
	sum := 0
	left := 0

	for left < len(s) {
		if left+1 < len(s) {
			val, ok := m[string(s[left])+string(s[left+1])]
			if ok {
				sum += val
				left += 2
				continue
			}
		}

		val := m[string(s[left])]
		sum += val
		left++

	}
	return sum
}
