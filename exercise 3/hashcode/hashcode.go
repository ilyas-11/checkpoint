package main

import (
	"fmt"
)

func main() {
	fmt.Println(HashCode("A"))
	fmt.Println(HashCode("AB"))
	fmt.Println(HashCode("BAC"))
	fmt.Println(HashCode("Hello World"))
}

func HashCode(dec string) string {
	str := ""
	for _, x := range dec {
		s := (int(x) + len(dec)) % 127
		if s < 33 {
			s += 33
		}
		str += string(rune(s))
	}
	return str
}
