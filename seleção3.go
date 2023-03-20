package main

import "fmt"

func main() {
	var numero int
	fmt.Println("O numero é", numero)
	fmt.Scan(&numero)

	if numero%2 == 0 {
		fmt.Println("Seu numero é par")
	} else {
		fmt.Println("Seu numero é impar")
	}
}
