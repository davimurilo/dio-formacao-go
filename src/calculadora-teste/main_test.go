package main

import "testing"

func TestSomar(teste *testing.T) {
	valor1 := 10.0
	valor2 := 2.0
	valorEsperado := 12.0
	resultado := Somar(valor1, valor2)

	VerificarSeFalhou(resultado, valorEsperado, teste)
}

func TestSomaNegativo(teste *testing.T) {
	valor1 := -10.0
	valor2 := -2.0
	valorEsperado := -12.0
	resultado := Somar(valor1, valor2)

	VerificarSeFalhou(resultado, valorEsperado, teste)
}

func TestSubtrair(teste *testing.T) {
	valor1 := 10.0
	valor2 := 2.0
	valorEsperado := 8.0
	resultado := Subtrair(valor1, valor2)

	VerificarSeFalhou(resultado, valorEsperado, teste)
}

func TestMultiplicar(teste *testing.T) {
	valor1 := 10.0
	valor2 := 2.0
	valorEsperado := 20.0
	resultado := Multiplicar(valor1, valor2)

	VerificarSeFalhou(resultado, valorEsperado, teste)
}

func TestDividir(teste *testing.T) {
	valor1 := 10.0
	valor2 := 2.0
	valorEsperado := 5.0
	resultado := Dividir(valor1, valor2)

	VerificarSeFalhou(resultado, valorEsperado, teste)
}

func VerificarSeFalhou(resultado, valorEsperado float64, teste *testing.T) {
	if resultado != valorEsperado {
		teste.Errorf("Falhou valor esperado era %g mas retornou %g", valorEsperado, resultado)
	}
}
