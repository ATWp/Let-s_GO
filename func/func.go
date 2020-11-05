package main

import "fmt"

func main() {
	var a = 29
	var b = 1

	var sum_ab int
	var sum_short_ab int

	sum_ab = summ(a, b)
	sum_short_ab = summ_short(a, b)
	fmt.Println("sum_ab", sum_ab)
	fmt.Println("sum_short_ab", sum_short_ab)
	//fmt.Println("summ_double_back", summ_double_back(a, b)) функцию в функцию неззя :((
	//эти переменные еще и описывать надо :((((((  коротка запись : не поможет :(
	var summ_double_back_a int
	var summ_double_back_b int
	summ_double_back_a, summ_double_back_b = summ_double_back(a, b) //нельзя выводить сразу список :(
	//var summ_double_back_mas [2]int
	//summ_double_back_mas = summ_double_back(a, b) записать два новых значения в массив созданный для этого не выйдет... круть!
	fmt.Println("summ_double_back_ab", summ_double_back_a, summ_double_back_b)
	//fmt.Println("summ_double_back_mas", summ_double_back_mas)
}

func summ(num_1 int, num_2 int) int {
	var res int
	res = num_1 + num_2
	return res
}

func summ_short(num_1 int, num_2 int) int {
	return num_1 + num_2
}

func summ_double_back(num_1 int, num_2 int) (int, int) {
	return num_1 + num_2, num_1 * num_2
}
