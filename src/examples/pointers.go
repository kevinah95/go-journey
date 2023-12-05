package examples

import "fmt"

func pointers() {
	s := "Hola"

	p := &s
	fmt.Println(p)
	fmt.Println(*p)
	*p = "¡Hola, Gophers!"
	fmt.Println(s)
	p = new(string)
	fmt.Println(p, *p)
}
