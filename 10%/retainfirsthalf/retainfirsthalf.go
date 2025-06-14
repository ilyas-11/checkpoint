package main

import (
	"fmt"
)

func main() {
	fmt.Println(RetainFirstHalf("This is the 1st halfThis is the 2nd half"))
	fmt.Println(RetainFirstHalf("A"))
	fmt.Println(RetainFirstHalf(""))
	fmt.Println(RetainFirstHalf("Hello World"))
}

func RetainFirstHalf(str string) string {
	s := ""
	if len(str) == 1 {
		return str
	}
	for i := 0; i < len(str)/2; i++ {
		s += string(str[i])
	}
	return s
}
