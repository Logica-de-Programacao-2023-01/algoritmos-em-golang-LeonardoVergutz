package main

import "fmt"

func main() {
	var (
		idade float64
	)
	fmt.Print("Qual é a sua idade em anos?")
	fmt.Scan(&idade)

	idadedias := idade * 365

	fmt.Println("Sua idade em dias é", idadedias)
}
