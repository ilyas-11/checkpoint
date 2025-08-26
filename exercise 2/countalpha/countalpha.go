package main

import "fmt"

func CountAlpha(s string) int {
	count := 0
	for _, x := range s {
		if x >= 'a' && x <= 'z' || x >= 'A' && x <= 'Z' {
			count++
		}

	}
	return count
}

func main() {
	fmt.Println(CountAlpha("Hello world"))
	fmt.Println(CountAlpha("H e l l o"))
	fmt.Println(CountAlpha("H1e2l3l4o"))
}
