package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(comma("12345678"))
	fmt.Println(commaWByte("12345678"))
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
