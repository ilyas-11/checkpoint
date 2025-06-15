package main

import (
	"fmt"
)

func main() {
	fmt.Println(Itoa(12345))
	fmt.Println(Itoa(0))
	fmt.Println(Itoa(-1234))
	fmt.Println(Itoa(987654321))
}
func Itoa(n int) string {
	str := ""
	signe := ""
	if n < 0 {
		signe = "-"
		n = -n
	}
	if n == 0 {
		return "0"
	}
	for n > 0 {
		str = string((n%10)+'0') + str
		n = n / 10
	}
	return signe + str
}
