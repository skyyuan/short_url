package utils

import (
	"strconv"
	"strings"
)

var (
	Tokens string
	Length int
)

func init() {
	// 0-9
	for i := 0; i <= 9; i++ {
		Tokens += strconv.Itoa(i)
	}

	// a-z
	for i := 0; i < 26; i++ {
		Tokens += string(byte('a') + byte(i))
	}

	// a-z
	for i := 0; i < 26; i++ {
		Tokens += string(byte('A') + byte(i))
	}

	Length = len(Tokens)
}

func IdToString(id int) string {
	var result string
	for id > 0 {
		d := id % Length
		result = string(Tokens[d]) + result
		id /= Length
	}
	return result
}

func StringToId(str string) int {
	var result = 0
	for _, s := range str {
		value := strings.Index(Tokens, string(s))
		result = result*Length + value
	}
	return result
}
