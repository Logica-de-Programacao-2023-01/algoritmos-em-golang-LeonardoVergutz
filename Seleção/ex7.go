package main

import "fmt"

func main() {
	var (
		salario float64
	)
	fmt.Print("Qual é o seu salario?")
	fmt.Scan(&salario)

	reajustemenor := (salario * 110) / 100
	reajuastemaior := (salario * 105) / 100

	if salario < 1000 {
		fmt.Println("O seu novo salário é", reajustemenor)
	} else if salario >= 1000 {
		fmt.Println("O seu novo salário é", reajuastemaior)
	}
}
