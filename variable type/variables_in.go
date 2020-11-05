package main

import "fmt"

func main() {
	/*Чтобы описать любую переменную используйте var <имя переменной> <тип переменной> = <значение>
	тип переменной, можно не указывать!
	А еще вы можете встретить такую интересную ошибку:
	# command-line-arguments
	.\variables_in.go:10:6: age_int declared but not use
	и единственное решение, которое предлагают: be a python developer, it's tough thing to use go-lint
	С переменными есть такая проблема в Go, как...
	*/
	var age = 12
	var num int = 13

	fmt.Println(age)
}
