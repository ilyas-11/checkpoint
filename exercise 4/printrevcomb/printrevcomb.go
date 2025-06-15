package main

func main() {
	//987, 986, 985, 984, 983, 982, 981, 980, 976, ..., 310, 210
	for a := 9; a >= 0; a-- {
		for b := a - 1; b >= 0; b-- {
			for c := b - 1; c >= 0; c-- {
				print(a)
				print(b)
				print(c)
				if a != 2 {
					print(",")
					print(" ")
				}
			}
		}
	}
}
