package main

import "fmt"

func main() {
	var (
		numero int
	)

	fmt.Print("Vamos descobrir se um número é ímpar ou par.")
	fmt.Print("Qual o número vc deseja verificar?")
	fmt.Scan(&numero)

	if numero%2 == 0 {
		fmt.Println("O numero é par")
	} else {
		fmt.Println("O numero é ímpar")
	}
}
