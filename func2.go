// Crie uma função que receba uma string e retorne a quantidade de vogais presentes.
package main

import (
	"fmt"
	"strings"
)

func stringVog(s string) (int, error) {
	vogais := "aeiou"
	i := 0

	for _, letras := range s {
		if strings.ContainsRune(vogais, letras) {
			i++
		}
	}
	return i, nil
}

func main() {
	s := "Pindamonhangaba"
	fmt.Println(stringVog(s))
}
