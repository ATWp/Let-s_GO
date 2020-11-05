package main

import "fmt"

func main() {
	//замыкания? lambda! прям мем представился пхпххпхп
	var num int = 3
	multiple := func(num int) int {
		num *= 3
		return num
	}
	//fmt.Println(multiple()) так он будет выводить память, где находится замыкание? объект?
	var res int
	res = multiple(num)
	fmt.Println(res)
}
