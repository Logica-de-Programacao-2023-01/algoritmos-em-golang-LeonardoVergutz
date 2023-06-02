package main

import "fmt"

func main() {
	var salario float64

	fmt.Print("Digite o seu salario:")
	fmt.Scan(&salario)

	novosalario := salario - ((salario * 10) / 100)

	fmt.Println("O seu salário com o desconto é:", novosalario)
}
