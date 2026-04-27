package main

import "fmt"

func main() {

	fmt.Println("Defasio de múltiplo de 3, de um intervalo de 1 a 100")
	fmt.Printf("\n")

	numeroMultiploDeTres := make([]int, 0)

	for i := 1; i <= 100; i++ {

		if i%3 == 0 {
			numeroMultiploDeTres = append(numeroMultiploDeTres, i)
		}
	}

	fmt.Printf("Os números múltiplos de 3 são : %v", numeroMultiploDeTres)

	fmt.Printf("\n\n")

	fmt.Println("Defasio de múltiplo de 3 e 5 (3 para Pin, 5 para Pan, 3 e 5 para Pin Pan), de um intervalo de 1 a 100")

	fmt.Printf("\n")

	for i := 1; i <= 100; i++ {

		switch {

		case (i%3 == 0) && (i%5 == 0):
			fmt.Println("Pin Pan")
		case (i%3 == 0):
			fmt.Println("Pin")
		case (i%5 == 0):
			fmt.Println("Pan")
		default:
			fmt.Println(i)
		}
	}

}
