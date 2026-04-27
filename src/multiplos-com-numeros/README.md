# Multiplos com Numeros

Este programa em Go resolve dois desafios no intervalo de 1 a 100.

## Objetivo

1. Listar os numeros multiplos de 3.
2. Exibir uma variacao do desafio tipo FizzBuzz:
   - Multiplo de 3: Pin
   - Multiplo de 5: Pan
   - Multiplo de 3 e 5: Pin Pan
   - Caso contrario: mostra o proprio numero

## Como o programa funciona

1. Mostra uma mensagem inicial do primeiro desafio.
2. Percorre os numeros de 1 a 100 e guarda os multiplos de 3 em um slice.
3. Imprime o slice com todos os multiplos encontrados.
4. Mostra uma mensagem inicial do segundo desafio.
5. Percorre novamente de 1 a 100 usando `switch` com condicoes:
   - Primeiro testa se o numero e multiplo de 3 e 5.
   - Depois testa se e multiplo de 3.
   - Depois testa se e multiplo de 5.
   - Se nao for nenhum dos casos, imprime o numero.

## Exemplo de comportamento

- Primeiro bloco: imprime os multiplos de 3 entre 1 e 100.
- Segundo bloco: imprime linhas com Pin, Pan, Pin Pan ou o numero, conforme as regras.

## Como executar

Na pasta do projeto, execute:

```bash
go run main.go
```

## Arquivo principal

- `main.go`