package main

import "fmt"

func main() {
	s := make([]int, 5, 8)
	s = append(s, 1, 2, 3, 4)
	printSlice(s)

	s2 := s[2:5]
	printSlice(s2)

	//s2[3] = 30
	printSlice(s2[:cap(s2)])

	s3 := make([]int, 2)
	copy(s3, s[1:3])
	printSlice(s3)
}

func printSlice(s []int) {
	fmt.Printf("Len is %d, cap is %d, value is %v\n", len(s), cap(s), s)
}
