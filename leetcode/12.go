package main

import "fmt"

func main() {
	fmt.Println(intToRoman(4))

}

func intToRoman(num int) string {
	a := num / 1 % 10
	b := num / 10 % 10
	c := num / 100 % 10
	d := num / 1000 % 10
	ret := ""
	if d > 0 {
		for i := 0; i < d; i++ {
			ret += "M"
		}
	}
	if c > 0 {
		if c <= 3 {
			for i := 0; i < c; i++ {
				ret += "C"
			}
		} else if c == 4 {
			ret += "CD"
		} else if c == 9 {
			ret += "CM"
		} else {
			s := "D"
			for i := 0; i < c-5; i++ {
				s += "C"
			}
			ret += s
		}
	}
	if b > 0 {
		if b <= 3 {
			for i := 0; i < b; i++ {
				ret += "X"
			}
		} else if b == 4 {
			ret += "XL"
		} else if b == 9 {
			ret += "XC"
		} else {
			s := "L"
			for i := 0; i < b-5; i++ {
				s += "X"
			}
			ret += s
		}
	}
	if a > 0 {
		if a <= 3 {
			for i := 0; i < a; i++ {
				ret += "I"
			}
		} else if a == 4 {
			ret += "IV"
		} else if a == 9 {
			ret += "IX"
		} else {
			s := "V"
			for i := 0; i < a-5; i++ {
				s += "I"
			}
			ret += s
		}
	}
	return ret
}
