package main

import (
	"fmt"
	"os"
	"bufio"
)

type Contato struct {
	Nome  string
	Email string
}

func (c *Contato) AlterarEmail(novoEmail string) {
	c.Email = novoEmail
}

func adicionaContato(c Contato, contatos [5]Contato) [5]Contato {
	for i := 0; i < len(contatos); i++ {
		if contatos[i].Nome == "" {
			contatos[i] = c
			break
		}
	}
	return contatos
}

func excluiContato(contatos [5]Contato) [5]Contato {
	for i := len(contatos) - 1; i >= 0; i-- {
		if contatos[i].Nome != "" {
			contatos[i] = Contato{}
			break
		}
	}
	return contatos
}

func main() {
	var contatos [5]Contato
	leitor := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Escolha uma opção:")
		fmt.Println("1 - Adicionar contato")
		fmt.Println("2 - Excluir último contato")

		var opcao int
		fmt.Scanln(&opcao)

		switch opcao {
		case 1:
			var c Contato
			fmt.Println("Informe o nome do contato:")
			c.Nome, _ = leitor.ReadString('\n')
			c.Nome = c.Nome[:len(c.Nome)-1]

			fmt.Println("Informe o e-mail do contato:")
			c.Email, _ = leitor.ReadString('\n')
			c.Email = c.Email[:len(c.Email)-1]

			contatos = adicionaContato(c, contatos)
		case 2:
			contatos = excluiContato(contatos)
		default:
			fmt.Println("Opção inválida!")
		}
	}
}
