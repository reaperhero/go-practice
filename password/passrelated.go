package password

import (
	"fmt"
	"math/rand"
	"regexp"
	"time"
)

const (
	numStr      = "0123456789"
	lowCharStr  = "abcdefghijklmnopqrstuvwxyz"
	upCharStr   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	allCharsStr = upCharStr + lowCharStr + numStr
	minLevel    = 3
)

// 生成特定长度的随机密码口令
func GeneratePassword(l int) (password string, err error) {
	if l < 8 || l > 20 {
		return "", fmt.Errorf("err passowrd lenth")
	}
	result := make([]byte, l)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < l; i++ {
		result[i] = allCharsStr[rand.Intn(len(allCharsStr))]
	}

	return string(result), nil
}

// checkPassword:检测密码合法性
func CheckPassword(password string) (ok bool, err error) {
	// 检测密码长度,至少8个字符长度，不超过20个字符
	if len(password) < 8 || len(password) > 20 {
		return false, fmt.Errorf("err passowrd lenth")
	}
	// 检测密码字符合法性,不能包含空格和中文字符(双字节字符)
	var re *regexp.Regexp
	re, _ = regexp.Compile(`[^0-9a-zA-Z\!\@\#\$\^\%\&\*\(\)]+`)
	if re.MatchString(password) {
		return false, fmt.Errorf("err password illegal")
	}
	// 检测密码强度
	level := 0
	patternList := []string{`[0-9]+`, `[a-z]+`, `[A-Z]+`}
	for _, pattern := range patternList {
		re, _ = regexp.Compile(pattern)
		if re.MatchString(password) {
			level++
		}
	}
	if level < minLevel {
		return false, fmt.Errorf("err password rule")
	}
	return true, nil
}
