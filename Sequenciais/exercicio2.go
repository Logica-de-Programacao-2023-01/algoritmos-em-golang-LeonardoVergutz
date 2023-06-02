package main

import "fmt"

func main() {
	var num int
	fmt.Print("Digite um número:")
	fmt.Scan(&num)

	dobro := num * 2
	triplo := num * 3
	quadruplo := num * 4

	fmt.Print("O dobro do numero é:", dobro, "\nO triplo do número é:", triplo, "\nO quadruplo do número é:", quadruplo)
}
