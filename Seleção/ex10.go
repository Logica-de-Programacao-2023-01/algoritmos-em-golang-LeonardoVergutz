package main

import (
	"fmt"
)

func main() {

	var idade int
	fmt.Print("Digite a sua idade:")
	fmt.Scan(&idade)

	if idade < 10 {
		fmt.Print("Você é considerado mirim!")
	} else if idade > 9 && idade < 14 {
		fmt.Print("Você é considerado infantil!")
	} else if idade >= 14 && idade < 18 {
		fmt.Print("Você é considerado juvenil!")
	} else {
		fmt.Print("Você é considerado adulto!")
	}
}
