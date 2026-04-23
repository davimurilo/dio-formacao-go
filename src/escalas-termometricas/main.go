package main

import "fmt"

func main() {

	temperaturaKelvin := 100.00
	temperaturaKelvinParaCelsius := temperaturaKelvin - 273.15

	fmt.Printf("A temperatura de %.2f K (Kelvin) é igual a %.2f °C (Celsius)", temperaturaKelvin, temperaturaKelvinParaCelsius)
}
