package main

import "fmt"

func main() {
	var (
		primeiro, segundo float64
	)
	fmt.Print("Qual é o primeiro número?")
	fmt.Scan(&primeiro)
	fmt.Print("Qual é o segundo número?")
	fmt.Scan(&segundo)

	multiplicação := primeiro * segundo
	soma := (primeiro) + (segundo)
	if primeiro > 0 && segundo > 0 {
		fmt.Println("O resultado é", multiplicação)
	} else {
		fmt.Println("O resultado é ", soma)
	}
}
