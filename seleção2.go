package main

import "fmt"

func main() {
	x := 7
	y := 2
	z := 4

	if x < y && x < z {
		fmt.Println("x é o menor número")
	} else if y < x && y < z {
		fmt.Println("y é o menor número")
	} else if z < x && z < y {
		fmt.Println("z é o menor número")
	}

}
