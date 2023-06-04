package main

import (
	"fmt"
)

func main() {
	a := 34
	b := 42.0

	fmt.Printf("%v\n", a*int(b)) // We can't multiply different types so we need to convert one of them
	fmt.Printf("%v of type %T\n", a, a)
	fmt.Printf("%v of type %T\n", b, b)

	//Conversion
	c := int(b)
	fmt.Printf("%v of type %T\n", c, c)
	b = float64(a)
	fmt.Printf("%v of type %T\n", b, b)

}
