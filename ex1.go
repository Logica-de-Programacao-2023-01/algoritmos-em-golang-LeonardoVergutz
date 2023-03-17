package main

import "fmt"

func main() {
	var (
		primeiro, segundo, terceiro float64
	)
	fmt.Print("Qual é o primeiro numero?")
	fmt.Scan(&primeiro)
	fmt.Print("Qual é o segundo numero?")
	fmt.Scan(&segundo)
	fmt.Print("Qual é o terceiro numero?")
	fmt.Scan(&terceiro)

	soma := primeiro + segundo + terceiro

	fmt.Println("O resultado é", soma)
}
