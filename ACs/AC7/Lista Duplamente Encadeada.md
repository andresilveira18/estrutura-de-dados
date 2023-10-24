# Lista Duplamente Encadeada

Uma lista duplamente encadeada é uma estrutura de dados em que cada nó possui referências tanto para o próximo nó quanto para o nó anterior na sequência. Isso facilita a navegação em ambas as direções e pode tornar certas operações mais eficientes.

## **Estrutura de um Nó**

Cada nó em uma lista duplamente encadeada contém três campos:

1. **Anterior**: Um ponteiro para o nó anterior na lista.
2. **Dados**: O valor ou informação armazenada no nó.
3. **Próximo**: Um ponteiro para o próximo nó na lista.

## **Operações Básicas**

### **1. Exibição dos Nós em uma Lista Duplamente Encadeada**

Esta operação percorre a lista do primeiro ao último nó, imprimindo o valor armazenado em cada nó. Podemos começar do nó cabeça e seguir os ponteiros "próximo" até alcançar o final da lista (quando o ponteiro "próximo" é **`NULL`**).

- Definir o nó atual como a cabeça da lista.
- Enquanto o nó atual não for **`NULL`**, fazer:
    - Imprimir o valor do nó atual.
    - Mover para o próximo nó (nó atual = nó atual.próximo).

```
Função ExibirListaDupla(nó):
    Se nó é NULL então
        Escrever "A lista está vazia"
        Retornar

    nóAtual <- nó
    Enquanto nóAtual ≠ NULL Fazer
        Escrever nóAtual.dados
        nóAtual <- nóAtual.próximo

Fim Função
```

### **2. Busca de um Nó**

A operação de busca percorre a lista até encontrar um nó com o valor desejado.

- Definir o nó atual como a cabeça da lista.
- Enquanto o nó atual não for **`NULL`**, fazer:
    - Verificar se o nó atual tem o valor desejado.
        - Se tiver, retorne o nó atual.
    - Mover para o próximo nó.
- Se o final da lista for alcançado sem encontrar o valor, retornar **`NULL`** ou uma indicação de que o valor não foi encontrado.

```
Função BuscarNó(nó, valor):
    Se nó é NULL então
        Escrever "A lista está vazia"
        Retornar NULL

    nóAtual <- nó
    Enquanto nóAtual ≠ NULL Fazer
        Se nóAtual.dados = valor então
            Retornar nóAtual

        nóAtual <- nóAtual.próximo

    Escrever "Valor não encontrado"
    Retornar NULL
Fim Função
```

### **3. Inserção de um Nó**

A inserção pode ocorrer em várias posições: no início, no meio ou no fim da lista.

    ```
    Função InserirNóOrdenado(cabeça, valor):
        novoNó <- criar novo Nó
        novoNó.dados <- valor
    
        Se cabeça = NULL então
            //  A lista está vazia, o novo nó se torna a cabeça.
            cabeça <- novoNó
            Retornar
    
        Se valor < cabeça.dados então
            //  O novo nó deve ser inserido antes da cabeça atual.
            novoNó.próximo <- cabeça
            cabeça.anterior <- novoNó
            cabeça <- novoNó
    
        Senão
            //  Percorra a lista para encontrar o local adequado para a inserção.
            nó_atual <- cabeça
            Enquanto nó_atual.próximo ≠ NULL e valor > nó_atual.dados
                nó_atual <- nó_atual.próximo
            
            Se valor > nó_atual.dados então
                // Inserir no final da lista.
                nó_atual.próximo <- novoNó
                novoNó.anterior <- nó_atual
    
            Senão
                // Inserir no meio da lista.
                novoNó.próximo <- nó_atual
                novoNó.anterior <- nó_atual.anterior
                nó_atual.anterior.próximo <- novoNó
                nó_atual.anterior <- novoNó
    
        Retornar
    Fim Função
    ```

### **4. Remoção de um Nó**

A remoção de um nó requer ajustes nos ponteiros para remover o nó da lista.

- Localizar o nó a ser removido.
- Se o nó a ser removido é a cabeça da lista:
    - Atualizar a cabeça para ser o próximo nó.
    - Fazer o "anterior" da nova cabeça (se não for **`NULL`**) apontar para **`NULL`**.
- Se o nó a ser removido não é a cabeça:
    - Fazer o "próximo" do nó anterior apontar para o próximo do nó a ser removido.
    - Se o próximo do nó a ser removido não for **`NULL`**, fazer o "anterior" dele apontar para o nó anterior do nó a ser removido.

```
Função RemoverNó(cabeça, nóARemover):
    Se cabeça = NULL OU nóARemover = NULL então
        Escrever "Operação inválida"
        Retornar cabeça

    Se cabeça = nóARemover então
        cabeça <- nóARemover.próximo

    Se nóARemover.próximo ≠ NULL então
        nóARemover.próximo.anterior <- nóARemover.anterior

    Se nóARemover.anterior ≠ NULL então
        nóARemover.anterior.próximo <- nóARemover.próximo

    Retornar cabeça
Fim Função
```
