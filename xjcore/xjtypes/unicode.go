package xjtypes

import "unicode"

//是否数字
func IsDigit(str string) bool {
	for _, x := range []rune(str) {
		if !unicode.IsDigit(x) {
			return false
		}
	}
	return true
}

//是否电子邮箱地址
func IsEmail(str string) bool {
	for _, x := range []rune(str) {
		if x == '@' {
			return true
		}
	}
	return false
}
