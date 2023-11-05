package main

import "fmt"

type No struct {
	valor int
	esq   *No
	dir   *No
}

type ArvoreBinaria struct {
	raiz *No
}

func busca(valor int, no *No) (int, *No) {
	if no == nil {
		return 0, nil
	}
	if valor == no.valor {
		return 1, no
	}
	if valor < no.valor {
		if no.esq == nil {
			return 2, no
		}
		return busca(valor, no.esq)
	} else {
		if no.dir == nil {
			return 3, no
		}
		return busca(valor, no.dir)
	}
}

func (arvore *ArvoreBinaria) insere(valor int) string {
	if arvore.raiz == nil {
		arvore.raiz = &No{valor: valor}
		return "Valor inserido como raiz"
	}
	f, no := busca(valor, arvore.raiz)
	switch f {
	case 1:
		return "Inserção inválida"
	case 2:
		no.esq = &No{valor: valor}
	case 3:
		no.dir = &No{valor: valor}
	}
	return "Valor inserido"
}

func main() {
	arvore := ArvoreBinaria{}

	fmt.Println(arvore.insere(10))
	fmt.Println(arvore.insere(5))
	fmt.Println(arvore.insere(15))
	fmt.Println(arvore.insere(10)) 
	fmt.Println(arvore.insere(6))
	fmt.Println(arvore.insere(1))
	fmt.Println(arvore.insere(20))

	_, no := busca(15, arvore.raiz)
	if no != nil && no.valor == 15 {
		fmt.Println("Valor 15 encontrado")
	} else {
		fmt.Println("Valor 15 não encontrado")
	}

}
	