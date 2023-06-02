package main

import "fmt"

func main() {
	var num, soma, quantidade int

	for {
		fmt.Print("Digite um número:")
		fmt.Scan(&num)
		if num == 0 {
			break
		}

		soma += num
		quantidade += 1

	}
	if quantidade > 0 {
		media := soma / quantidade
		fmt.Println("A media dos valores informados é:", float64(media))
	} else {
		fmt.Println("Nenhum valor foi informado!")
	}

}
