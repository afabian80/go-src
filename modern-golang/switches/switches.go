package main

import "fmt"

func main() {
	x := 2

	switch x {
	case 2:
		fmt.Println("x is 2")
		//fallthrough
	case 3:
		fmt.Println("x is 3")
	case 4:
		fmt.Println("x is 4")
	default:
		fmt.Println("x is unknown")
	}

	if x = 3; x == 3 {
		fmt.Println("x was 2 indeed")
	}
}
