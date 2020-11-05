package main

import "fmt"

func main() {
	var age = 20
	switch age {
	case 5:
		fmt.Println("Вам 5 лет!")
	case 10:
		fmt.Println("Вам 10 лет!")
	case 15:
		fmt.Println("Вам 15 лет!")
	default:
		fmt.Println("Вам неизвестно сколько лет!")
	}
}
