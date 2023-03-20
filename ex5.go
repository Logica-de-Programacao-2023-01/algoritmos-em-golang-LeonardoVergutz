package main

import "fmt"

func main() {
	var (
		numero int
	)

	fmt.Print("Vamos saber se um número é múltiplo de 3 e 5 ao mesmo tempo.")
	fmt.Print("Qual o numero vc deseja saber?")
	fmt.Scan(&numero)

	if numero%3 == 0 && numero%5 == 0 {
		fmt.Println("O numero é multiplo de 3 e 5 ao mesmo tempo")
	} else {
		fmt.Println("O numero não é multiplo de 3 e 5 ao mesmo tempo")
	}

}
