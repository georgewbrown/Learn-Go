package main

import "fmt"

// name := "Carl" // this will not work because short declaration does not work outside fucntion scope
var name = "Carl"

func main() {
	fmt.Println(name)
}
