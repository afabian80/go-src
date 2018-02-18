package main

import "fmt"

func main() {
	m1 := make(map[string]int)
	m1["egy"] = 1
	v, ok := m1["harom"]
	fmt.Println(v, ok)

	fmt.Println(m1)
	delete(m1, "ketto")
	fmt.Println(m1)
}
