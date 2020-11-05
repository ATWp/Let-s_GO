package main

import "fmt"

func main() {
	var age = 20
	if age < 5 {
		fmt.Println("Вам пора в детский сад!")
	} else if age == 5 {
		fmt.Println("Вам пора в preschool!")
	} else if (age > 5) && (age <= 18) {
		var grade = age - 5
		fmt.Println("Вам пора в ", grade, "school!")
	} else {
		fmt.Println("Вам пора в universe!")
	}
}
