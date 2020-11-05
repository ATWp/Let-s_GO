package main

import "fmt"

func main() {
	var web string = "test"
	fmt.Println(web + " it's cool")

	var num float64 = 4.324124
	fmt.Printf("%f \n", num)
	fmt.Printf("%.2f \n", num) // 2 символа после запятой
	fmt.Printf("%T \n", num)   // тип переменной

	var isDone bool = true
	fmt.Printf("%t \n", isDone)
}
