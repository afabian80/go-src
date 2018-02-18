package main

import "fmt"

type person struct {
	Name    string
	Age     int
	Address string
}

func main() {
	akos := person{
		Name:    "Akos",
		Age:     37,
		Address: "Hungary",
	}

	fmt.Println(akos)
	fmt.Println(akos.Name)
}
