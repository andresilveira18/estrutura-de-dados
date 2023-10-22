# **Lista Circular**

Uma lista circular simplesmente encadeada é uma estrutura de dados em que cada nó aponta para o próximo nó na lista, e o último nó aponta de volta para o primeiro nó, formando um círculo. Diferentemente de uma lista encadeada simples, ela não possui um fim claro, permitindo uma iteração contínua pela lista.

## **Estrutura de um Nó**

Cada nó em uma lista circular simplesmente encadeada contém dois campos:

1. **Dados**: O valor ou informação armazenada no nó.
2. **Próximo**: Um ponteiro para o próximo nó na lista.

## **Operações Básicas**

### **1. Exibição dos Nós em uma Lista Circular**

A operação de exibição percorre a lista do início ao fim, imprimindo o valor de cada nó. Uma peculiaridade das listas circulares é que o último nó aponta de volta para o primeiro, criando um ciclo. Isso significa que, para percorrer a lista, você deve começar no primeiro nó e continuar até chegar novamente a esse nó. A função para exibir os nós pode ser assim:

- Comece com o nó cabeça.
- Imprima o valor do nó atual.
- Vá para o próximo nó.
- Repita os passos 2 e 3 até que o nó atual seja novamente o nó cabeça.
    
    ```
    Função ExibirListaCircular(nó):
        Se nó é NULL então
            Escrever "A lista está vazia"
            Retornar
        Fim Se
    
        nóAtual <- nó
        Fazer
            Escrever nóAtual.dados
            nóAtual <- nóAtual.próximo
        Enquanto nóAtual ≠ nó
    Fim Função
    ```
    

### **2. Inserção de um Nó no Início da Lista**

A inserção no início de uma lista circular envolve a criação de um novo nó e o ajuste dos ponteiros para incluir esse nó na lista:

- Crie um novo nó.
- Faça o próximo do novo nó apontar para a cabeça atual da lista.
- Encontre o último nó da lista (o nó cujo próximo aponta para a cabeça).
- Faça o próximo do último nó apontar para o novo nó.
- Atualize a cabeça da lista para ser o novo nó.
    
    ```
    Função InserirNoInicio(nó, valor):
        novoNó <- criar novo Nó
        novoNó.dados <- valor
        
        Se nó é NULL então
            novoNó.próximo <- novoNó  // A lista estava vazia, então o novo nó aponta para si mesmo
        Senão
            nóAtual <- nó
            Enquanto nóAtual.próximo ≠ nó Fazer
                nóAtual <- nóAtual.próximo
            Fim Enquanto
            
            nóAtual.próximo <- novoNó  // O último nó na lista agora aponta para o novo nó
            novoNó.próximo <- nó  // O novo nó aponta para o que era anteriormente o primeiro nó
        Fim Se
        
        Retornar novoNó  // O novo nó agora é o primeiro nó na lista
    Fim Função
    ```
    

### **3. Exclusão do Primeiro Nó da Lista**

A exclusão do primeiro nó em uma lista circular requer ajustes nos ponteiros para remover o nó da lista:

- Verifique se a lista está vazia. Se estiver, a operação não pode ser concluída.
- Se a lista tem apenas um nó (a cabeça aponta para si mesma), simplesmente defina a cabeça como NULL.
- Encontre o último nó da lista (o nó cujo próximo aponta para a cabeça).
- Faça o próximo do último nó apontar para o segundo nó da lista (o próximo da cabeça atual).
- Libere a memória do nó cabeça.
- Atualize a cabeça da lista para ser o segundo nó.
    
    ```
    Função ExcluirPrimeiroNó(nó):
        Se nó é NULL então
            Escrever "A lista está vazia"
            Retornar NULL
        Fim Se
        
        Se nó.próximo é nó então
            // A lista tem apenas um nó
            Liberar nó
            Retornar NULL
        Senão
            nóAtual <- nó
            Enquanto nóAtual.próximo ≠ nó Fazer
                nóAtual <- nóAtual.próximo
            Fim Enquanto
            
            nóAtual.próximo <- nó.próximo  // O último nó agora aponta para o segundo nó
            Liberar nó
            Retornar nóAtual.próximo  // O segundo nó agora é o primeiro
        Fim Se
    Fim Função
    ```
