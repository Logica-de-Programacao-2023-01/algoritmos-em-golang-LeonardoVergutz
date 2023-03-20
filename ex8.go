package main

import "fmt"

func main() {
	var (
		numero int
	)

	fmt.Print("Digite um número de 1 a 7 para saber seu dia da semana correspondente.")
	fmt.Scan(&numero)

	switch numero {
	case 1:
		fmt.Println("Domingo")
	case 2:
		fmt.Println("Segunda")
	case 3:
		fmt.Println("Terça")
	case 4:
		fmt.Println("Quarta")
	case 5:
		fmt.Println("Quinta")
	case 6:
		fmt.Println("Sexta")
	case 7:
		fmt.Println("Sabado")
	default:
		fmt.Println("Número inválido! Digite um número entre 1 e 7.")

	}

}
