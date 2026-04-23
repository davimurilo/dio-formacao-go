package main

import "fmt"

func main() {

	temperaturaCelsius := 100.0
	temperaturaFahrenheit := 100.0

	temperaturaCelsiusParaFahrenheit := ((9 * temperaturaCelsius) / 5) + 32
	temperaturaFahrenheitParaCelsius := ((temperaturaFahrenheit - 32) * 5) / 9

	fmt.Printf("A temperatura %f em Celsius, será %f em Fahrenheit.", temperaturaCelsius, temperaturaCelsiusParaFahrenheit)
	fmt.Println()
	fmt.Printf("A temperatura %f em Fahrenheit, será %f em Celsius.", temperaturaFahrenheit, temperaturaFahrenheitParaCelsius)
	fmt.Println()
}
