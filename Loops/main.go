package main

import "fmt"

// a sintxe de for em go é muito parecida com a de C, mas com algumas diferenças
func main() {
	var res int
	for i := 0; i < 10; i++ {
		res++
	}
	fmt.Println(res) // output: 10

	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for range arr {
		// vai iterar sobre todos os itens do array e vai realizar uma ação em cada indice
	}

	for i, elem := range arr {
		// i é o indice do array em cada posição e elem é o elemento em cada posição do array no indice i
		fmt.Println(i, elem)
	}

	for _, elem := range arr {
		fmt.Println(elem) // dessa forma o índice é ignorado e a ação só iram acontecer nos elementos do array
	}

	for range 10 {
		// é a mesma coisa que escrever " for i := 0; i < 10; i++", porém de um jeito simplificado
	}

	// vale lembrar que escrevendo dessa forma eu só posso iterar sobre um elemento
	for i := range 10 {
		fmt.Println(i) // output: 0 até 9
	}

	arr2 := [3]int{1, 2, 3}

	for i, elem := range arr2 {
		fmt.Println(&i, &elem) // output: endereço de memória de cada índice e elemento, mas vale destacar que cada indice e cada elemento é recriado a cada iteração
	}
}
