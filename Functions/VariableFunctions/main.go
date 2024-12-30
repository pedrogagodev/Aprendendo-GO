package main

import "fmt"

func main() {
	fmt.Println(sum(10, 10, 10))
}

func sum(nums ...int) int { // os 3 pontos indicam que a função pode receber um numero indefinido de argumentos
	var out int
	for _, n := range nums {
		out += n
	}
	return out
}
