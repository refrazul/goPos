package main

import "fmt"
import "palina/gpos/iso/interpreter"

func main() {
	cadena := []byte{49, 57, 78}
	s := string(interpreter.Interpret(cadena, 0))

	fmt.Println("Cadena: " + s)
}
