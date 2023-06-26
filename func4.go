package main

import (
	"fmt"
)

func segundoMaior(slice []int) int {
	var maiorNum int
	var segundoMaior int
	for i := 0; i < len(slice); i++ {
		if slice[i] > maiorNum {
			segundoMaior = maiorNum
			maiorNum = slice[i]
		}
	}
	return segundoMaior
}

func main() {
	slice := []int{10, 5, 8, 20, 7}
	fmt.Print(segundoMaior(slice))
}
