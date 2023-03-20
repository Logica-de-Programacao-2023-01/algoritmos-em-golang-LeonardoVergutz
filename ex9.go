package main

import "fmt"

func main() {
	var (
		primeiro, segundo, terceiro float64
	)

	fmt.Print("Escreva o primeiro numero.")
	fmt.Scan(&primeiro)
	fmt.Print("Escreva o segundo numero.")
	fmt.Scan(&segundo)
	fmt.Print("Escreva o terceiro numero.")
	fmt.Scan(&terceiro)

	if primeiro > segundo && segundo > terceiro {
		fmt.Println("Em ordem crescente os numeros ficam assim", terceiro, segundo, primeiro)
	} if primeiro < segundo


}
