package main

import (
	"fmt"
	"strings"
)

func decode(cipherText, keyword string) string {
	var res []int32

	for idx, symbol := range cipherText {
		kIndex := idx % len(keyword)
		currentShift := keyword[kIndex] - 'A'
		symbol = ((symbol-'A')-int32(currentShift)+26)%26 + 'A'
		res = append(res, symbol)
	}

	return (string(res))
}

func encode(message, keyword string) string {
	message = strings.ToUpper(strings.Replace(message, " ", "", -1))
	var res []rune

	for idx, symbol := range message {
		kIndex := idx % len(keyword)
		currentShift := keyword[kIndex] - 'A'
		symbol := ((symbol-'A')+int32(currentShift))%26 + 'A'
		res = append(res, symbol)
	}
	return string(res)
}

func main() {
	message := "Meet me at midnight"
	key := "GOLANG"

	cipher := encode(message, key)
	fmt.Println("Coded message: ", cipher)

	fmt.Println("Decoded message: ", decode(cipher, key))
}
