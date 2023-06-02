package main

import "fmt"

func main() {
	var peso, altura float64

	fmt.Print("Digite a sua altura:")
	fmt.Scan(&altura)
	fmt.Print("Digite o seu peso:")
	fmt.Scan(&peso)

	imc := peso / (altura * altura)

	fmt.Printf("O seu IMC é:%.2f", imc)
}
