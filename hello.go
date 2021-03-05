package main

import (
	"first-go-project/mycalculator"
	"fmt"
)

func main() {
	fmt.Printf("Hello world!\n")
	resultado := mycalculator.Add(10, 5)
	fmt.Printf("%v\n", resultado)
	fmt.Printf("%T\n", resultado)
	resultadoX := mycalculator.AddX(10)
	fmt.Printf("%v\n", resultadoX)
	fmt.Printf("%T\n", resultadoX)
	fmt.Printf("%v", mycalculator.A)
}
