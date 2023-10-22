# .md 2

Uma lista duplamente encadeada é uma estrutura de dados em que cada nó possui referências tanto para o próximo nó quanto para o nó anterior na sequência. Isso facilita a navegação em ambas as direções e pode tornar certas operações mais eficientes.

## **Estrutura de um Nó**

Cada nó em uma lista duplamente encadeada contém três campos:

1. **Anterior**: Um ponteiro para o nó anterior na lista.
2. **Dados**: O valor ou informação armazenada no nó.
3. **Próximo**: Um ponteiro para o próximo nó na lista.

## **Operações Básicas**

### **1. Exibição dos Nós em uma Lista Duplamente Encadeada**

Esta operação percorre a lista do primeiro ao último nó, imprimindo o valor armazenado em cada nó. Você pode começar do nó cabeça e seguir os ponteiros "próximo" até alcançar o final da lista (quando o ponteiro "próximo" é **`NULL`**).

- Defina o nó atual como a cabeça da lista.
- Enquanto o nó atual não for **`NULL`**, faça:
    - Imprima o valor do nó atual.
    - Mova para o próximo nó (nó atual = nó atual.próximo).

```
Função ExibirListaDupla(nó):
    Se nó é NULL então
        Escrever "A lista está vazia"
        Retornar
    Fim Se

    nóAtual <- nó
    Enquanto nóAtual ≠ NULL Fazer
        Escrever nóAtual.dados
        nóAtual <- nóAtual.próximo
    Fim Enquanto
Fim Função
```

### **2. Busca de um Nó**

A operação de busca percorre a lista até encontrar um nó com o valor desejado.

- Defina o nó atual como a cabeça da lista.
- Enquanto o nó atual não for **`NULL`**, faça:
    - Verifique se o nó atual tem o valor desejado.
        - Se tiver, retorne o nó atual.
    - Mova para o próximo nó.
- Se o final da lista for alcançado sem encontrar o valor, retorne **`NULL`** ou uma indicação de que o valor não foi encontrado.

```
Função BuscarNó(nó, valor):
    Se nó é NULL então
        Escrever "A lista está vazia"
        Retornar NULL
    Fim Se

    nóAtual <- nó
    Enquanto nóAtual ≠ NULL Fazer
        Se nóAtual.dados = valor então
            Retornar nóAtual
        Fim Se
        nóAtual <- nóAtual.próximo
    Fim Enquanto

    Escrever "Valor não encontrado"
    Retornar NULL
Fim Função
```

### **3. Inserção de um Nó**

A inserção pode ocorrer em várias posições: no início, no meio ou no fim da lista. Aqui, vou descrever a inserção no início da lista.

- Crie um novo nó.
- Faça o "próximo" do novo nó apontar para a cabeça atual da lista.
- Faça o "anterior" da cabeça atual da lista apontar para o novo nó.
- Atualize a cabeça da lista para ser o novo nó.

```
Função InserirNó(nó, valor):
    novoNó <- criar novo Nó
    novoNó.dados <- valor
    novoNó.próximo <- nó
    novoNó.anterior <- NULL
    
    Se nó ≠ NULL então
        nó.anterior <- novoNó
    Fim Se

    Retornar novoNó
Fim Função
```

### **4. Remoção de um Nó**

A remoção de um nó requer ajustes nos ponteiros para remover o nó da lista.

- Localize o nó a ser removido.
- Se o nó a ser removido é a cabeça da lista:
    - Atualize a cabeça para ser o próximo nó.
    - Faça o "anterior" da nova cabeça (se não for **`NULL`**) apontar para **`NULL`**.
- Se o nó a ser removido não é a cabeça:
    - Faça o "próximo" do nó anterior apontar para o próximo do nó a ser removido.
    - Se o próximo do nó a ser removido não for **`NULL`**, faça o "anterior" dele apontar para o nó anterior do nó a ser removido.
- Libere a memória do nó a ser removido.

```
Função RemoverNó(cabeça, nóARemover):
    Se cabeça = NULL OU nóARemover = NULL então
        Escrever "Operação inválida"
        Retornar cabeça
    Fim Se

    Se cabeça = nóARemover então
        cabeça <- nóARemover.próximo
    Fim Se

    Se nóARemover.próximo ≠ NULL então
        nóARemover.próximo.anterior <- nóARemover.anterior
    Fim Se

    Se nóARemover.anterior ≠ NULL então
        nóARemover.anterior.próximo <- nóARemover.próximo
    Fim Se

    Liberar nóARemover
    Retornar cabeça
Fim Função
```