package main

import (
	"fmt"
	"strings"
)

func main() {
	palavra := "Ô James, eu quero uma salada de frutas!"
	fmt.Println(separarEspaco(palavra))
}

func separarEspaco(palavra string) ([]string, error) {
	palavras := []string{}
	if palavra == "" {
		return palavras, fmt.Errorf("String vazia.")
	}

	separador := strings.Split(palavra, " ")

	return separador, nil
}
