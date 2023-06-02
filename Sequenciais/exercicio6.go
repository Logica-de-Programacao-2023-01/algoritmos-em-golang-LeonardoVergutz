package main

import "fmt"

func main() {
	var salario float64

	fmt.Print("Digite o seu salario:")
	fmt.Scan(&salario)

	novosalario := salario + (salario * 0.15)

	fmt.Print("O seu salario após o aumento de 15% é:", novosalario)
}
