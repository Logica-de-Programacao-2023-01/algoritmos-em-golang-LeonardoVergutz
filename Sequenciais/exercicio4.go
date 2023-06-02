package main

import "fmt"

func main() {
	var num1, num2, num3 int

	fmt.Print("Digite um número:")
	fmt.Scan(&num1)
	fmt.Print("Digite um número:")
	fmt.Scan(&num2)
	fmt.Print("Digite um número:")
	fmt.Scan(&num3)

	media := float64((num1*2)+(num2*3)+(num3*5)) / 10.

	fmt.Println("A media ponderada é:", media)
}
