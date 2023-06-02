package main

import "fmt"

func main() {
	var dias float64
	var diaria float64

	fmt.Print("Quantos dias voce trabalhou?")
	fmt.Scan(&dias)
	fmt.Print("Qual é o valor da sua diaria?")
	fmt.Scan(&diaria)

	salario := dias * diaria

	fmt.Println("O seu salário é:", salario)

}
