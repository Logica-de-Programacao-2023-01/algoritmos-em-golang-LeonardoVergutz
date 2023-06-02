package main

import "fmt"

func main() {
	var (
		peso float64
	)
	fmt.Print("Qual é o seu peso em quilos?")
	fmt.Scan(&peso)

	libras := peso * 2.2046
	fmt.Println("O seu peso em libras é", libras)

}
