package main

import "fmt"

func main() {
	//откладывание - выполняется, после того, как сработала функция main
	defer one()
	two()
}

func one() {
	fmt.Println(" func one ")
}

func two() {
	fmt.Println(" func two ")
}
