// Crie uma função que receba um slice de inteiros e retorne a média aritmética dos valores.
package main

import "fmt"

func mediaAritmetica(slice []int) int {
	soma := 0
	for _, valor := range slice {
		soma += valor
	}
	return soma / len(slice)
}

func main() {
	slice := []int{10, 5, 12}
	fmt.Println(mediaAritmetica(slice))
}
