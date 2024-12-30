package main

import "fmt"

func main() {
	res, rem := divide(10, 3)
	fmt.Println(res, rem)
}


/* posso declarar a minha variável no segundo bloco de parametros da função ao inves de declarar e atribuir um valor dentro da função
   com a variável já declarada nesse segundo bloco de argumentos, eu posso utilizar somente "return" */

/* vale lembrar que muito dificilmente esse tipo de retorno vai ser utilizado, somente com funções muito pequenas onde fica muito claro 
   oq ta sendo retornado, mas na maioria dos casos o certo é especificar o que esta sendo retornado */
func divide(a, b int) (res int, rem int) {
	res = a / b
	rem = a%b
	return
}