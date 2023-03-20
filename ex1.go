package main

import "fmt"

func main() {
	var (
		primeiro, segundo float64
	)
	fmt.Print("Qual é o primeiro numero?")
	fmt.Scan(&primeiro)
	fmt.Print("Qual é o segundo numero?")
	fmt.Scan(&segundo)

	if primeiro > segundo {
		fmt.Println(" O maior número é", primeiro)
	} else {
		fmt.Println("O maior numero é", segundo)
	}

}
