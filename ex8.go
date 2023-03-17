package main

import "fmt"

func main() {
	var (
		dias, valor float64
	)
	fmt.Print("Quantos dias voce trabalhou?")
	fmt.Scan(&dias)
	fmt.Print("Quanto vc recebeu por dia?")
	fmt.Scan(&valor)

	salario := dias * valor
	fmt.Println("Seu salario é", salario)
}
