package main

import "fmt"

// Estrutura para um nó da lista
type No struct {
	Dado int
	Proximo *No
}

// Estrutura para a lista circular
type ListaCircular struct {
	Cabeca *No
}

// Função para criar uma nova lista circular
func NovaListaCircular() *ListaCircular {
	return &ListaCircular{}
}

// Função para adicionar um nó no início da lista
func (lista *ListaCircular) AdicionarNoInicio(dado int) {
	novoNo := &No{Dado: dado}
	if lista.Cabeca == nil {
		// A lista está vazia, então o novo nó aponta para si mesmo
		novoNo.Proximo = novoNo
	} else {
		// Encontrar o último nó da lista
		ultimoNo := lista.Cabeca
		for ultimoNo.Proximo != lista.Cabeca {
			ultimoNo = ultimoNo.Proximo
		}
		// Adicionar o novo nó no início
		novoNo.Proximo = lista.Cabeca
		ultimoNo.Proximo = novoNo
	}
	// Atualizar a cabeça da lista para o novo nó
	lista.Cabeca = novoNo
}

// Função para remover o primeiro nó da lista
func (lista *ListaCircular) RemoverPrimeiroNo() {
	if lista.Cabeca == nil {
		fmt.Println("A lista está vazia")
		return
	}
	if lista.Cabeca.Proximo == lista.Cabeca {
		// A lista tem apenas um nó
		lista.Cabeca = nil
	} else {
		// Encontrar o último nó da lista
		ultimoNo := lista.Cabeca
		for ultimoNo.Proximo != lista.Cabeca {
			ultimoNo = ultimoNo.Proximo
		}
		// Remover o primeiro nó
		ultimoNo.Proximo = lista.Cabeca.Proximo
		lista.Cabeca = lista.Cabeca.Proximo
	}
}

// Função para exibir todos os nós da lista
func (lista *ListaCircular) ExibirNos() {
	if lista.Cabeca == nil {
		fmt.Println("A lista está vazia")
		return
	}
	noAtual := lista.Cabeca
	for {
		fmt.Print(noAtual.Dado, " ")
		noAtual = noAtual.Proximo
		if noAtual == lista.Cabeca {
			break
		}
	}
	fmt.Println()
}

func main() {
	lista := NovaListaCircular()

	lista.AdicionarNoInicio(10)
	lista.AdicionarNoInicio(20)
	lista.AdicionarNoInicio(30)
	lista.AdicionarNoInicio(40)

	fmt.Print("Lista Circular: ")
	lista.ExibirNos()

	lista.RemoverPrimeiroNo()
	fmt.Print("Lista após remover o primeiro nó: ")
	lista.ExibirNos()
}