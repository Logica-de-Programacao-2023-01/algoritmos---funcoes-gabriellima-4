// Crie uma função que receba um slice de strings e retorne a concatenação de todas as strings.
package main

import "fmt"

func stringConcat(frase []string) (string, error) {
	var result string

	for _, i := range frase {
		result += string(i)
	}

	return result, nil
}

func main() {
	frase := []string{"Oi", ",", "tudo", " ", "bem?"}
	fmt.Println(stringConcat(frase))
}
