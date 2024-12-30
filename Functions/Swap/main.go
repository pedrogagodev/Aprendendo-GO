package main

import "fmt"

func main() {
    /* o operador " := " é usado para declarar e inicializar uma ou mais variáveis
	   o compilador de go consegue inferir o tipo da variável automaticamente de acordo com o valor que ela irá receber */
	a, b := swap(10, 20)
	fmt.Println(a, b)
}

func swap(a, b int) (int, int) {
	return b, a
}