package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(convert("PAYPALISHIRING", 3))
}

func convert(s string, numRows int) string {
	if numRows < 2 {
		return s
	}
	res := make([]string, numRows)
	i, flag := 0, -1
	for _, v := range s {
		res[i] += string(v)
		if i == 0 || i == numRows-1 {
			flag = -flag
		}
		fmt.Println(i, flag)
		i += flag
	}
	return strings.Join(res, "")
}
