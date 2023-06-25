// Escreva uma função que receba um slice de strings como parâmetro e
// retorne uma string com todas as strings concatenadas e separadas por vírgulas. Caso o slice esteja vazio, retorne um erro.
package main

import (
	"fmt"
	"strings"
)

func concatenarStringsVirgulas(sl []string) (string, error) {
	var result string
	if sl == nil {
		return "cade a frase po**a?", fmt.Errorf("O slice é vazio.")
	}
	for _, palavra := range sl {
		result += string(palavra)
	}
	return strings.ReplaceAll(result, " ", ","), nil
}

func main() {
	slice := []string{"Oi", " ", "tudo", "bem?"}
	fmt.Println(concatenarStringsVirgulas(slice))
}
