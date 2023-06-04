package main

import (
	"fmt"
)

func add(x string, y string) (string, string) {
	return y, x
}

func main() {

	fmt.Println(add("Hello", "World"))

}

func sayHello() {
	fmt.Println("Hello")
}
