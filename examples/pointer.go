package main

import (
	"fmt"
)

func main() {

	//	un tipo de variable especial que mantiene la dirección de otra variable en memoria
	hello := "Hello Word"
	pointer := &hello
	*pointer = "Hello Romel"

	fmt.Println(hello, pointer, *pointer)
}
