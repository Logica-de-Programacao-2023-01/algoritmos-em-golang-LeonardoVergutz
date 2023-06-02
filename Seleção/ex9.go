package main

import "fmt"

func main() {
	var numero1, numero2, numero3 int
	fmt.Print("Digite o primeiro número:")
	fmt.Scan(&numero1)
	fmt.Print("Digite o segundo número:")
	fmt.Scan(&numero2)
	fmt.Print("Digite o segundo número:")
	fmt.Scan(&numero3)

	if numero1 > numero2 && numero1 > numero3 {
		fmt.Print("O maior número é:", numero1)
	} else if numero2 > numero1 && numero2 > numero3 {
		if numero1 == numero3 {
			fmt.Print("O primeiro e o terceiro número são iguais!")
		} else {
			fmt.Print("O maior número é:", numero2)
		}
	} else if numero1 == numero3 && numero2 != numero3 {

	} else if numero3 > numero2 && numero3 > numero1 {
		fmt.Print("O maior número é:", numero3)
	} else if numero1 == numero2 && numero2 == numero3 {
		fmt.Print("Os tres números inseridos são iguais!")
	} else if numero1 == numero2 && numero2 != numero3 {
		fmt.Print("Os dois primeiros numeros são iguais!")
	} else if numero2 == numero3 && numero2 != numero1 {
		fmt.Print("Os dois ultimos números são iguais")
	}
}
