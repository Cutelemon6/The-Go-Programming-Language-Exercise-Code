package main

import (
	"fmt"
	"os"
	"strings"
)

//!+
func main() {
	fmt.Println(strings.Join(os.Args[0:], "-")) // args[0] 表示名称，运行后是为带路径的可执行程序名
}
