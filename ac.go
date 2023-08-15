package main

import (
	"fmt"
)

func main() {

	fmt.Println(fibo(7))

	fmt.Println(dia_da_semana(3))
	fmt.Println(dia_da_semana(9))

	fmt.Println(e_bissexto(1995))
	fmt.Println(e_bissexto(2012))
	fmt.Println(e_bissexto(1900))
	fmt.Println(e_bissexto(2000))
}

// Não consegui a 1 :(

func fibo(n int) int {
	if n <= 1 {
		return n
	}
	return fibo(n-1) + fibo(n-2)
}

func dia_da_semana(n int) string {
	switch n {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-feira"
	case 3:
		return "Terça-feira"
	case 4:
		return "Quarta-feira"
	case 5:
		return "Quinta-feira"
	case 6:
		return "Sexta-feira"
	case 7:
		return "Sábado"
	default:
		return "Valor inválido"
	}
}

func e_bissexto(ano int) bool {
	if ano%400 == 0 {
		return true
	} else if ano%100 == 0 {
		return false
	} else if ano%4 == 0 {
		return true
	}
	return false
}
