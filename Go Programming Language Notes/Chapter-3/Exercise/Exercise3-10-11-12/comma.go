package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(comma("12345678"))
	fmt.Println(commaWByte("12345678"))
	fmt.Println(commaWSignFloat("-12345678.123456756"))
	fmt.Println(isAnagram("Burak", "Deniz"))
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

// Exercise 3.10
func commaWByte(s string) string {
	length := len(s)
	if length <= 3 {
		return s
	}
	var buf bytes.Buffer
	reminder, times := length%3, length/3

	if reminder != 0 {
		buf.WriteString(s[:reminder])
		buf.WriteString(",")
	}

	for i := 0; i < times; i++ {
		buf.WriteString(s[i*3+reminder : i*3+reminder+3])
		if i != times-1 {
			buf.WriteString(",")
		}
	}
	return buf.String()
}

// Exercise 3.11
func commaWSignFloat(s string) string {
	var sign string
	var numberString string
	if s[0:1] == "+" || s[0:1] == "-" {
		sign = s[0:1]
		numberString = s[1:]
	}
	var buf bytes.Buffer
	buf.WriteString(sign)
	periodIndex := strings.LastIndex(s, ".")
	if periodIndex != -1 {
		buf.WriteString(commaWByte(numberString[:periodIndex-1]))
		buf.WriteString(".")
		buf.WriteString(numberString[periodIndex+1:])
	} else {
		buf.WriteString(commaWByte(numberString))
	}
	return buf.String()
}

// Exercise 3.12
func isAnagram(firstWord, secondWord string) bool {
	m := make(map[rune]int)

	for _, char := range firstWord {
		m[char]++
	}
	for _, char := range secondWord {
		if i, ok := m[char]; ok {
			if i > 1 {
				m[char]--
			} else {
				delete(m, char)
			}
		} else {
			return false
		}
	}
	return len(m) == 0
}
