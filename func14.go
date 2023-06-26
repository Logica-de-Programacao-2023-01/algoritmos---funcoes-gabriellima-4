package main

import (
	"fmt"
)

func main() {
	slice1 := []int{1, 3, 15, 37, 50}
	slice2 := []int{2, 5, 9, 80}
	fmt.Print(numerosDoisSlices(slice1, slice2))
}
func numerosDoisSlices(slice1 []int, slice2 []int) ([]int, error) {
	slice3 := []int{}
	if len(slice1) == 0 || len(slice2) == 0 {
		return slice3, fmt.Errorf("Slice vazio")
	}
	for _, s1 := range slice1 {
		for _, s2 := range slice2 {
			if s1 == s2 {
				slice3 = append(slice3, s1)
			}
		}
	}
	return slice3, nil
}
