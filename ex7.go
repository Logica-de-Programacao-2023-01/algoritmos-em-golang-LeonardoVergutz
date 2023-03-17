package main

import "fmt"

func main() {
	var (
		numero float64
	)
	fmt.Print("Qual é o numero")
	fmt.Scan(&numero)

	antecessor := numero - 1
	sucessor := numero + 1

	fmt.Println("O antecessor do numero é", antecessor)
	fmt.Println("O sucessor do numero é", sucessor)
}
