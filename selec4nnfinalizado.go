package main

import "fmt"

func main() {

	var altura int
	fmt.Print("Qual sua altura?")
	fmt.Scan(&altura)
	fmt.Println("ok", altura)

	var sexo string
	fmt.Print("Qual o seu sexo?")
	fmt.Scan(&sexo)
	fmt.Println("ok", sexo)

	var peso int
	fmt.Print("Qual o seu peso?")
	fmt.Scan(&peso)
	fmt.Println("ok", peso)

}
