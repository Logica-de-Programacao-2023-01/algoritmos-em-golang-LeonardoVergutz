package main

import "fmt"

func main() {
	var num int
	fmt.Print("Digite um número:")
	fmt.Scan(&num)
	fmt.Print("A tabuada do número inserido é:")
	for i := -1; i <= 10; i++ {
		tabuada := i * num
		fmt.Println(tabuada)
	}

}
