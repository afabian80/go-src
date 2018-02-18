package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	as := [5]int{1, 2, 3, 4, 5}
	for a := range as {
		fmt.Println(a)
	}

	i := 0
	for i <= 3 {
		fmt.Println(i)
		i++
	}
}
