package main

import "fmt"

func main() {
	//указатели - позволяют передавать в функцию не просто параметр или переменную, а именно адрес переменной
	var x = 0 // тип переменной не должен быть указан!!
	pointer(x)
	fmt.Println("pointer", x)

	pointer_nice(&x)
	fmt.Println("pointer_nice", x)
}

func pointer(x int) {
	x = 2
}

func pointer_nice(x *int) {
	*x = 2
}
