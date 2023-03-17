package main

import "fmt"

func main() {
	var (
		peso, altura float64
	)

	fmt.Print("Qual é o seu peso?")
	fmt.Scan(&peso)
	fmt.Print("Qual é a sua altura?")
	fmt.Scan(&altura)

	IMC := peso / (altura * altura)

	fmt.Println("O seu IMC é", IMC)

}
