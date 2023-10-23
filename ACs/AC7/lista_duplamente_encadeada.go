package main

import "fmt"

type No struct {
	Dado     int
	Proximo  *No
	Anterior *No
}

type ListaDupla struct {
	Cabeca *No
}

func NovaListaDupla() *ListaDupla {
	return &ListaDupla{}
}

func (lista *ListaDupla) AdicionarNoInicio(dado int) {
	novoNo := &No{Dado: dado}
	novoNo.Proximo = lista.Cabeca
	novoNo.Anterior = nil

	if lista.Cabeca != nil {
		lista.Cabeca.Anterior = novoNo
	}
	lista.Cabeca = novoNo
}

func (lista *ListaDupla) ExibirNos() {
	noAtual := lista.Cabeca
	if noAtual == nil {
		fmt.Println("A lista está vazia.")
		return
	}
	for noAtual != nil {
		fmt.Print(noAtual.Dado, " ")
		noAtual = noAtual.Proximo
	}
	fmt.Println()
}

func (lista *ListaDupla) BuscarNo(valor int) *No {
	noAtual := lista.Cabeca
	for noAtual != nil {
		if noAtual.Dado == valor {
			return noAtual
		}
		noAtual = noAtual.Proximo
	}
	return nil
}

func (lista *ListaDupla) RemoverNo(no *No) {
	if no == nil {
		fmt.Println("O nó a ser removido não pode ser nulo.")
		return
	}

	if no.Anterior != nil {
		no.Anterior.Proximo = no.Proximo
	} else {
		lista.Cabeca = no.Proximo
	}

	if no.Proximo != nil {
		no.Proximo.Anterior = no.Anterior
	}
}

func main() {
	lista := NovaListaDupla()

	lista.AdicionarNoInicio(10)
	lista.AdicionarNoInicio(20)
	lista.AdicionarNoInicio(30)
	lista.AdicionarNoInicio(40)

	fmt.Print("Lista Dupla: ")
	lista.ExibirNos()

	no := lista.BuscarNo(20)
	if no != nil {
		fmt.Println("Nó encontrado:", no.Dado)
	} else {
		fmt.Println("Nó não encontrado.")
	}

	lista.RemoverNo(no)
	fmt.Print("Lista após remover o nó: ")
	lista.ExibirNos()
}