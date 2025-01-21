package main

import (
	"strconv"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	strx := strconv.Itoa(x)
	reverseStrx := reverse(strx)

	if strx == reverseStrx {
		return true
	}

	return false
}

// 文字列を逆順にする関数
func reverse(str string) string {
	var reverseStr string
	for i := len(str) - 1; i >= 0; i-- {
		reverseStr += string(str[i])
	}
	return reverseStr
}
