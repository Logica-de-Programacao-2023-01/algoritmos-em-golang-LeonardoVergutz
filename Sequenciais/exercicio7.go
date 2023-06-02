package main

import "fmt"

func main() {
	var numero int

	fmt.Print("Digite um número:")
	fmt.Scan(&numero)

	antecessor := numero - 1
	sucessor := numero + 1

	fmt.Print("O antecessor do número informado é:", antecessor, "\nO Sucessor do número informado é:", sucessor)
}
