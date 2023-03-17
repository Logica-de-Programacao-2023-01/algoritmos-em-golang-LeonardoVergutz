package main

import "fmt"

func main() {
	var (
		salario float64
	)
	fmt.Print("Qual é o seu salario?")
	fmt.Scan(&salario)
	reajuste := (salario * 115) / 100

	fmt.Println("Seu novo salario é", reajuste)
}
