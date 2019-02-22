package printer

import "fmt"

//Hello us an exported function
func Hello() {
	fmt.Println("unexported hello")
}
