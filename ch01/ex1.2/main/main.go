package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args[1:] {
		fmt.Println(i, ": "+arg) // i is an int and arg is a string
	}
}
