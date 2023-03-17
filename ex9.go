package main

import "fmt"

func main() {
	var (
		preço float64
	)
	fmt.Print("Qual é o preço do produto?")
	fmt.Scan(&preço)

	desconto := (preço * 90) / 100
	fmt.Println("O preço com desconto é de", desconto)
}
