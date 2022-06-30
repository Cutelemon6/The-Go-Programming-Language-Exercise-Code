package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	var s, sep string
	for _, arg := range os.Args[1:] {
		s += sep + arg
	}
	fmt.Println(s)
	fmt.Println("This process takes", time.Since(start).Seconds(), "seconds")

	start = time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println("this process takes", time.Since(start).Seconds(), "seconds")
}
