package main

import (
	"fmt"
)

func numeroPrimo(numero int) (bool, error) {
	primo := true
	if numero < 0 {
		return false, fmt.Errorf("Um número menor que zero não é primo.")
	}

	if numero < 2 {
		primo = false
	}

	for i := 2; i < numero; i++ {
		if numero%i == 0 {
			primo = false
		}
	}
	return primo, nil
}

func main() {
	numero := 83
	fmt.Print(numeroPrimo(numero))
}
