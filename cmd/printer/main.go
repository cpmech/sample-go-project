package main

import (
	"_/lib/count"
	"_/lib/list"
	"fmt"
)

func main() {
	for _, word := range list.Words("Hello World") {
		fmt.Println(word)
	}
	fmt.Printf("number of words = %d\n", count.Words("Hello World"))
}
