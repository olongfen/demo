package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestPalindrome("ibvjkmpyzsifuxcabqqpahjdeuzaybqsrsmbfplxycsafogotliyvhxjtkrbzqxlyfwujzhkdafhebvsdhkkdbhlhmaoxmbkqiwiusngkbdhlvxdyvnjrzvxmukvdfobzlmvnbnilnsyrgoygfdzjlymhprcpxsnxpcafctikxxybcusgjwmfklkffehbvlhvxfiddznwumxosomfbgxoruoqrhezgsgidgcfzbtdftjxeahriirqgxbhicoxavquhbkaomrroghdnfkknyigsluqebaqrtcwgmlnvmxoagisdmsokeznjsnwpxygjjptvyjjkbmkxvlivinmpnpxgmmorkasebngirckqcawgevljplkkgextudqaodwqmfljljhrujoerycoojwwgtklypicgkyaboqjfivbeqdlonxeidgxsyzugkntoevwfuxovazcyayvwbcqswzhytlmtmrtwpikgacnpkbwgfmpavzyjoxughwhvlsxsgttbcyrlkaarngeoaldsdtjncivhcfsaohmdhgbwkuemcembmlwbwquxfaiukoqvzmgoeppieztdacvwngbkcxknbytvztodbfnjhbtwpjlzuajnlzfmmujhcggpdcwdquutdiubgcvnxvgspmfumeqrofewynizvynavjzkbpkuxxvkjujectdyfwygnfsukvzflcuxxzvxzravzznpxttduajhbsyiywpqunnarabcroljwcbdydagachbobkcvudkoddldaucwruobfylfhyvjuynjrosxczgjwudpxaqwnboxgxybnngxxhibesiaxkicinikzzmonftqkcudlzfzutplbycejmkpxcygsafzkgudy"))
}
func longestPalindrome(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}
	max_len := 1
	left := 0
	dp := make([][]bool, n)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, n)
		dp[i][i] = true
	}
	for l := 2; l < n+1; l++ {
		for i := 0; i < n; i++ {
			j := l + i - 1
			if j >= n {
				break
			}
			if s[i] != s[j] {
				dp[i][j] = false
			} else {
				if j-i < 3 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}
			if dp[i][j] && j-i+1 > max_len {
				max_len = j - i + 1
				left = i
			}
		}
	}

	return s[left : left+max_len]

}

func longestPalindrome1(s string) string {
	if len(s) == 0 {
		return ""
	}
	if len(s) == 1 {
		return s
	}
	left := 0
	right := left
	arr := ""
	maxArr := ""
	fc := func(m1, m2 string) string {
		if len(m1) >= len(m2) {
			return m1
		}
		return m2
	}
	for left < len(s)-1 && right < len(s) {
		arr += string(s[right])
		if isPalindrome(arr) {
			right++
			maxArr = fc(maxArr, arr)
		} else if right == len(s)-1 {
			left++
			right = left
			arr = ""
		} else {
			right++
		}
	}
	return fc(maxArr, arr)

}

// 反转字符串
func isPalindrome(s string) bool {
	// 双指针法判断回文字符串
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}

	return true
}
