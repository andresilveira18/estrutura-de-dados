package main

import "fmt"

func par(array []int, alvo int) (int, int) {
	for i, j := 0, len(array)-1; i < j; {
		if s := array[i] + array[j]; s == alvo {
			return array[i], array[j]
		} else if s < alvo {
			i++
		} else {
			j--
		}
	}
	return -1, -1
}

func main() {
	array, alvo := []int{2, 3, 4, 5, 6, 7, 8}, 7
	if a, b := par(array, alvo); a != -1 {
		fmt.Printf("Par encontrado: (%d, %d)\n", a, b)
	} else {
		fmt.Println("Nenhum par encontrado.")
	}
}