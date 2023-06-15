// Crie uma função que receba um slice de inteiros e um valor inteiro,
// e retorne a posição do primeiro elemento igual ao valor no slice. Caso não encontre, retorne -1.
package main

import "fmt"

func valorIGelemento(slice []int, valor int) int {
	menosUm := -1
	for i := 0; i < len(slice); i++ {
		if slice[i] == valor {
			menosUm = i
		}
	}
	return menosUm
}

func main() {
	slice := []int{2, 5, 8, 14}
	num := 8
	result := valorIGelemento(slice, num)
	fmt.Println(result)
}
