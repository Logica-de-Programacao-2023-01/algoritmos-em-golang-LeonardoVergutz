package main

import "fmt"

func main() {
	var (
		primeiro, segundo, terceiro float64
	)

	fmt.Print("Qual é o primeiro número?")
	fmt.Scan(&primeiro)
	fmt.Print("Qual é o segundo numero?")
	fmt.Scan(&segundo)
	fmt.Print("Qual é o terceiro numero?")
	fmt.Scan(&terceiro)

	if primeiro < segundo && primeiro < terceiro {
		fmt.Println("O menor número é", primeiro)
	}
	if segundo < primeiro && segundo < terceiro {
		fmt.Println("O menor numero é ", segundo)
	}
	if terceiro < primeiro && terceiro < segundo {
		fmt.Println("O menor numero é", terceiro)
	}
}
