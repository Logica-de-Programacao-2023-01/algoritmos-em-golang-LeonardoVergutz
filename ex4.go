package main

import "fmt"

func main() {
	var (
		altura, peso, sexo float64
	)

	fmt.Print("Qual é o seu peso?")
	fmt.Scan(&peso)
	fmt.Print("Qual é a sua altura?")
	fmt.Scan(&altura)
	fmt.Print("Qual é o seu sexo?")
	fmt.Scan(&sexo)

	IMC := peso / (altura * altura)

	if IMC >= 18.5 && IMC <= 24.9 {
		fmt.Println("Você está dentro do peso ideal, parabéns")
	}
	if IMC < 18.5 {
		fmt.Println("Você está abaixo do peso ideal")
	}
	if IMC > 24.9 {
		fmt.Println("Você está acima do peso ideal")
	}
}
