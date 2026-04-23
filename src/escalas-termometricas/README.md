# Escalas Termométricas

Este programa em Go converte uma temperatura informada em Kelvin para Celsius e exibe o resultado no terminal.

## Objetivo

Demonstrar uma conversão simples entre escalas termométricas usando operações matemáticas básicas e saída formatada com `fmt.Printf`.

## Como funciona

O programa define um valor fixo em Kelvin:

- `100.00 K`

Em seguida, aplica a fórmula de conversão:

```text
°C = K - 273.15
```

Depois disso, o resultado é exibido no console com duas casas decimais.

## Exemplo de saída

```text
A temperatura de 100.00 K (Kelvin) é igual a -173.15 °C (Celsius)
```

## Como executar

No terminal, dentro da pasta do projeto, execute:

```bash
go run main.go
```

## Arquivo principal

O código-fonte do programa está em `main.go`.