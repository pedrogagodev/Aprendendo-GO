package main

import (
	"fmt"
	"math"
)

// condicionais em go são muito parecidas com as condicionais das outras linguagens, com o adicional de poder inicializar uma variável no inicio da condicional

func main() {
	if x := math.Sqrt(4); x < 10 {
		fmt.Println("Sqrt < 10")
	}
}
