package main

import "fmt"

func main() {

	var altura float64
	fmt.Print("Qual sua altura?")
	fmt.Scan(&altura)
	fmt.Println("ok", altura)

	var sexo string
	fmt.Print("Qual o seu sexo?")
	fmt.Scan(&sexo)
	fmt.Println("ok", sexo)

	var peso float64
	fmt.Print("Qual o seu peso?")
	fmt.Scan(&peso)
	fmt.Println("ok", peso)

	IMC := peso / (altura * altura)

	var class string
	if IMC <= 18.5 {
		class = "Abaixo do peso normal"
		println(class)
	} else if IMC >= 18.5 && IMC <= 24.9 {
		class = "Peso normal"
		println(class)
	} else if IMC >= 25.0 && IMC <= 29.9 {
		class = "Excesso de peso"
		println(class)
	} else if IMC >= 30.0 && IMC <= 34.9 {
		class = "Obesidade GRAU 1"
		println(class)
	} else if IMC >= 35.0 && IMC <= 39.9 {
		class = "Obesidade GRAU 2"
		println(class)
	} else {
		class = "Obesidade GRAU 3"
	}

}
