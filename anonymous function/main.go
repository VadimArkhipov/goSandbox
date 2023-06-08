package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	countSum := func(array []int) int {
		sum := 0
		for _, elem := range array {
			sum += elem
		}
		return sum
	}

	sum := countSum(a)
	fmt.Println(sum)

}
