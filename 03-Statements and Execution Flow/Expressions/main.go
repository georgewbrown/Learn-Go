package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.NumCPU() + 1)
}

// func main() {
// 	fmt.Println("Hello!" + "!!")
// }
