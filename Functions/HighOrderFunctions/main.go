package main

import "fmt"

func main() {
	x := sum(2)(3)
	/* basicamente o que foi feito foi uma função composta, onde eu tenho:
	   f := somar(2)
	   x := f(3)
	   ou seja, f passa o argumento da primeira função (a) como "2"
	   e o x passa o argumento da segunda função (b) como 3
	*/
	fmt.Println(x)
}

//a função sum irá receber um agumento 'a' inteiro e ira retornar uma função que vai receber um paramtro inteiro e que vai retornar um inteiro
func sum(a int) func(int) int {  
	return func(b int) int {
		return a + b
	}
}
