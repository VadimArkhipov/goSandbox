package main

import "fmt"

func main() {
	message := "L fdph, L vdz, L frqtxhuhg"
	var decodeMessage []rune

	for _, symbol := range []rune(message) {
		if 'A' <= symbol && 'Z' >= symbol {
			symbol -= 3
			if symbol < 'A' {
				symbol += 26
			}
		} else if 'a' <= symbol && 'z' >= symbol {
			symbol -= 3
			if symbol < 'a' {
				symbol += 26
			}
		}
		decodeMessage = append(decodeMessage, symbol)
	}

	fmt.Println(string(decodeMessage))
}
