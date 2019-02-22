package main

import "fmt"
import f "fmt"

func main() {
	fmt.Println("Hello Gopher!", 1, 12)
	f.Println("this works")
	bye()
	hey()
}

///// func main() prints:
//		Hello Gopher! 1 12
//		Bye!
//		Hey!
