/ Escreva uma função que receba uma string como parâmetro e
// retorne uma nova string com todas as vogais substituídas por "*". Caso a string seja vazia, retorne um erro.
package main

import (
	"fmt"
	"strings"
)

func vogaisAsteriscos(s string) (string, error) {

	s = strings.ReplaceAll(s, "a", "*")
	s = strings.ReplaceAll(s, "e", "*")
	s = strings.ReplaceAll(s, "i", "*")
	s = strings.ReplaceAll(s, "o", "*")
	s = strings.ReplaceAll(s, "u", "*")
	
	if s == "" {
		return "", fmt.Errorf("A string é vazia.")
	}
	return s, nil
}
