package main

import "fmt"

/* Arrays em go são parecidos com os arrays em C, os arrays em go são imutáveis no sentido de tamanho
e o tamanho dele tem que ser especificado quando ele foi criado */

func main() {
	arr := [3]int{}  // os colchetes vão indicar o tamanho do array, logo após virá o tipo dele e depois os valores (0 se nao especificado, ou "" se for string)
	fmt.Println(arr) // output: [0, 0, 0]

	const x = 10 // constantes podem ser atribuídas no tamanho do array, mas apenas constantes podem ser, e não variáveis
	arr2 := [x]string{5: "hello, world", 7: "foo"} // posso especificar previamente por indice os valores que eu quero usando essa sintaxe
	fmt.Println(arr2) // output: [       hello, world,     foo   ]
}