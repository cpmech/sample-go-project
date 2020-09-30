package main

import (
	"_/lib/count"
	"_/lib/list"
	"fmt"
)

func main() {
	words := list.Words("Hello World")
	nw := count.Words("Hello World")
	fmt.Printf("%v (%d)\n", words, nw)
}
