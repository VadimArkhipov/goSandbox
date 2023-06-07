package main

import (
	"fmt"
	"strings"
)

func sourceGopher(downstream chan string) {
	source := []string{"Controversy uh-huh", "Cause it's just another Saturday Night", "Hello"}
	for _, elem := range source {
		downstream <- elem
	}
	close(downstream)
}

func filterGopher(downstream, upstream chan string) {
	for elem := range upstream {
		words := strings.Fields(elem)
		for _, word := range words {
			downstream <- word
		}
	}
	close(downstream)
}

func printGopher(upstream chan string) {
	for elem := range upstream {
		fmt.Println(elem)
	}
}

func main() {
	c0 := make(chan string)
	c1 := make(chan string)

	go sourceGopher(c0)
	go filterGopher(c1, c0)
	printGopher(c1)
}
