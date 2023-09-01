package main

import "fmt"

var movimentos int

func hanoi(n int, origem string, auxiliar string, destino string) {
	if n == 1 {
		movimentos++
		fmt.Println("Mova o disco 1 de " + origem + " para " + destino)
		return
	}
	hanoi(n-1, origem, destino, auxiliar)
	movimentos++
	fmt.Println("Mova o disco " + fmt.Sprint(n) + " de " + origem + " para " + destino)
	hanoi(n-1, auxiliar, origem, destino)
}

func main() {
	var n int

	fmt.Println("Informe o número de discos:")
	fmt.Scanln(&n)

	hanoi(n, "A", "B", "C")

	fmt.Println("Número total de movimentos: " + fmt.Sprint(movimentos))
}