package main

import "fmt"

func main() {
	fmt.Println(sum(2, 3))
}


/* eu posso omitir algumas declarações de tipo caso todos os meu parâmetros sejam do mesmo tipo
   para fazer isso, basta eu colocar a tipagem após o último parametro */

func sum(a int, b int) int {
	return a + b
}