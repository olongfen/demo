package main

import "fmt"

func main() {
	fmt.Println(generateParenthesis(3))
}

func generateParenthesis(n int) []string {
	var (
		fc  = func(curStr string, left, right int) {}
		res []string
	)
	fc = func(curStr string, left, right int) {
		if left == 0 && right == 0 {
			res = append(res, curStr)
			return
		}
		if right < left {
			return
		}
		if left > 0 {
			fc(curStr+"(", left-1, right)
		}
		if right > 0 {
			fc(curStr+")", left, right-1)
		}
	}

	fc("", n, n)
	return res

}
