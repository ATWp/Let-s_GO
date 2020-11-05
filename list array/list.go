package main

import "fmt"

func main() {
	var arr [3]int
	arr[0] = 45
	arr[1] = 97
	arr[2] = 76
	fmt.Println(arr[1])

	nums := [4]float64{4.23, 5.23, 98.1, 66.43}
	for i, value := range nums {
		fmt.Println(i, value)
	}

}
