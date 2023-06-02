package main

import "fmt"

func main() {
	var idade int
	fmt.Print("Digite sua idade:")
	fmt.Scan(&idade)
	if idade > 120 {
		fmt.Print("Idade inserida invalida")
	}
	novaidade := idade * 365

	fmt.Println("A sua idade é aproximadamente", novaidade, "dias.")
}
