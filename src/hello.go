package main

import "fmt"

func main() {
	// Raw String Literal.
	rawLiteral := `あいう\n
  えお\n
  かきく`

	// Interpreted String Literal
	interLiteral := "あいうえお\nかきくけこ"

	fmt.Println(rawLiteral)
	fmt.Println()
	fmt.Println(interLiteral)
}
