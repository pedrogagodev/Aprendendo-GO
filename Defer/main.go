package main

import "fmt"

// defer é basicamente uma forma de assincronismo, uma função defer só vai ser executada depois que a função que está em volta em retornar algo
// porém os parametros passados para a função com defer serão avaliados imediatamente

// IMPORTANTE: defer obedece o escopo da função principal que está em volta,
// ou seja, se eu tenho um defer dentro de um if, por exemplo, e o if ta dentro de alguma outra função,
// o defer vai respeitar o escopo da funçao que está em volta do if, e nao do escopo do if

func main() {
	x := doDefer()
	fmt.Println(x)
	// output: hello <br> world <br> 10 pois o fluxo é o seguinte: ocorreu a chamada da função que está em volta do defer
	// > logo depois essa mesma função roda, mas sem executar a função defer dentro dela > a funão defer é chamada depois da "doDefer" retornar algo
	// > depois disso o fluxo vai para a função que chamou a "doDefer"
}

func doDefer() int {
	defer fmt.Println("world")
	fmt.Println("hello")
	return 10
}

func doDeferLIFO() {
	defer fmt.Println(3)
	defer fmt.Println(2)
	fmt.Println(1)

	// output: 1 2 3. o defer segue o padrao LIFO (last in, first out), ou seja, o ultimo defer a ser lido pelo compilador, é o primeiro a ser executado
	// em outras palavras: o primeiro defer a ser lido será o último a ser executado, vai seguir uma ordem de execuão de cima para baixo
}

func doDeferFuncAnoum() {
	x := 10
	defer func(y int) {
		fmt.Println(y)
	}(x)

	x = 50
	fmt.Println(x)

	// output: 50 10, pois o parametro que foi enviado para a função defer é avaliado imediatamente e a função defer "guarda" esse valor para usar depois

	z := 10
	defer func() {
		fmt.Println(z)
	}()

	z = 50
	fmt.Println(z)

	// output 50 50. pois o defer nao consegue avaliar a variavel z e guardar esse valor, ele só le depois que o z mudou de valor
}
