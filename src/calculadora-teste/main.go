package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

const (
	SOMA = iota + 1
	SUBTRACAO
	MULTIPLICAO
	DIVISAO
)

func main() {

	for {

		fmt.Println("##########################################")
		fmt.Println("# Programa Calculadora Operações Básicas #")
		fmt.Println("##########################################")

		fmt.Println()
		fmt.Println("Escolha o tipo de operação, digitando o número correspondente")
		fmt.Println("[1] Adição")
		fmt.Println("[2] Subtração")
		fmt.Println("[3] Multiplicação")
		fmt.Println("[4] Divisão")
		fmt.Println("[0] Encerrar")

		var erro error
		valor1 := 0.0
		valor2 := 0.0
		tipo := 0
		fmt.Print("Digite o tipo de operação:")

		_, erro = fmt.Scanln(&tipo)
		VerificarEncerramentoPorErro(erro)

		if tipo == 0 {
			os.Exit(1)
		}

		if tipo < 0 || tipo > 4 {
			fmt.Println("Tipo inválido deve ser 0,1,2,3 ou 4")
			os.Exit(1)
		}

		fmt.Print("Digite o valor 1:")
		_, erro = fmt.Scanln(&valor1)
		VerificarEncerramentoPorErro(erro)

		fmt.Print("Digite o valor 2:")
		_, erro = fmt.Scanln(&valor2)
		VerificarEncerramentoPorErro(erro)

		resultado := ExecutarOperacao(tipo, valor1, valor2)

		fmt.Printf("Resultado: %g \n", resultado)
		fmt.Println("Pressionar qualquer tecla para continuar...")
		fmt.Scanln()
		LimparConsole()
	}
}

func LimparConsole() {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}

func VerificarEncerramentoPorErro(err error) {
	if err != nil {
		fmt.Println("Erro ao capturar o tipo " + err.Error())
		os.Exit(1)
	}
}

func ExecutarOperacao(tipo int, valor1, valor2 float64) float64 {
	total := 0.0

	switch tipo {
	case SOMA:
		total = Somar(valor1, valor2)
	case SUBTRACAO:
		total = Subtrair(valor1, valor2)
	case MULTIPLICAO:
		total = Multiplicar(valor1, valor2)
	case DIVISAO:
		total = Dividir(valor1, valor2)
	}

	return total
}

func Somar(valor1, valor2 float64) float64 {
	return valor1 + valor2
}

func Subtrair(valor1, valor2 float64) float64 {
	return valor1 - valor2
}

func Multiplicar(valor1, valor2 float64) float64 {
	return valor1 * valor2
}

func Dividir(valor1, valor2 float64) float64 {
	return valor1 / valor2
}
