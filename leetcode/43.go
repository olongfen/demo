package main

import (
	"fmt"
	"strconv"
)

func multiply(num1 string, num2 string) string {
	n1, _ := strconv.ParseInt(num1, 10, 64)
	n2, _ := strconv.ParseInt(num2, 10, 64)
	return strconv.FormatInt(n1*n2, 10)
}

func main() {
	fmt.Println(int64(498828660196) * int64(840477629533))
}
