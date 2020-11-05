package main

import "fmt"

func main() {
	var i = 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	var i = 10
	for i > 0 {
		fmt.Println(i)
		i--
	}

	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}
}
