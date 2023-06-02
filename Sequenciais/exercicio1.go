package main

import "fmt"

func main() {
	var numero1, numero2, numero3 int

	fmt.Print("Digite o primeiro número:")
	fmt.Scan(&numero1)
	fmt.Print("Digite o segundo número:")
	fmt.Scan(&numero2)
	fmt.Print("Digite o terceiro número:")
	fmt.Scan(&numero3)

	soma := numero1 + numero2 + numero3

	fmt.Println("A soma dos número inseridos é:", soma)

}
