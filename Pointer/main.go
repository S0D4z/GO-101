package main

import "fmt"

func appendHello(str *string) {
	*str = "Hello, " + *str
}

func main() {
	m := 4
	// double(&m)
	fmt.Println(m)

	name := "Soda"
	appendHello(&name)
	fmt.Println(name)
}
