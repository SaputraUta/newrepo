package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Uta"
	names[1] = "Saputra"
	names[2] = "Depi"
	for i := 0; i < 3; i++ {
		fmt.Println(names[i])
	}
}
