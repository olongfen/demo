package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(myAtoi("--5"))

}

func myAtoi(s string) int {
	lastNum := false
	data := []byte{}
	for i := 0; i < len(s); i++ {
		if s[i] == 45 || s[i] == 43 {
			if !lastNum {
				data = append(data, s[i])
				lastNum = true
			} else {
				break
			}

		} else if s[i] >= 48 && s[i] <= 57 {
			data = append(data, s[i])
			lastNum = true
		} else if lastNum {
			break
		} else if s[i] != ' ' {
			break
		}
	}
	str := string(data)
	d, _ := strconv.Atoi(str)
	min := -1 << 31
	max := 1<<31 - 1
	if d < min {
		return min
	}
	if d > max {
		return max
	}
	return d
}
