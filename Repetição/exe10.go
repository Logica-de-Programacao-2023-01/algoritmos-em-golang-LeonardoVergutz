package main

import "fmt"

func main() {
	var num int
	maior := 0
	menor := 0

	for {
		fmt.Print("Digite um número:")
		fmt.Scan(&num)
		if num == 0 {
			break
		}
		if num > maior {
			maior = num
		}
		if num < menor {
			menor = num
		}
	}

	fmt.Println("O maior número informado é:", maior, "\nO menor número inforfmado é:", menor)

}
