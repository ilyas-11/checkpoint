package main

import "os"

//. "you   see   it's   easy   to   display   the   same   thing"
func main() {
	if len(os.Args) != 2 {
		return
	}
	c := ""
	str := []string{}
	s := os.Args[1]
	for i := 0; i < len(s); i++ {
		if s[0] == ' ' {
			s = s[1:]
			i--
		} else if s[i] == ' ' {
			str = append(str, c)
			c = ""
		} else if s[i] != ' ' {
			c += string(s[i])
		}

	}
	print(str)
	print(c)
}
