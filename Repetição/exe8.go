package main

import "fmt"

func main() {
	var num int

	fmt.Print("Digite um número:")
	fmt.Scan(&num)

	for i := 1; i <= 100000; i++ {
		if num%i == 0 {
			fmt.Println(i)
		}

	}
}
