package main

import "fmt"

// bool

// int  int8  int16  int32 	int64 (inteiros positivos e negativos)
// uint uint8 uint16 uint32 uint64 (inteiros positivos)

// byte (basicamente uma outra forma de se escrever uint8)

// rune (basicamente uma outra forma de se escrever int32)

// float32 float64

// complex64 complex128

func main() {
	const x = 5 // declaração de uma constante
	takeInt32(x) // usando a constante
	takeInt64(10) // constant literal
	takeString("foo") // string literal
}

func takeString(x string) {
	fmt.Println(x)
}

func takeInt32(x int32) {
	fmt.Println(x)
}

func takeInt64(x int64) {
	fmt.Println(x)
}