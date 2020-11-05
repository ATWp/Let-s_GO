package main

import "fmt"

type Cats struct {
	name     string
	age      int
	happines float64
}

/*А вот и методы для класса...Интересно, почему они не сделали все под одну структуру класса?
1. Возможно из-за удобства структуры без методов?
2. Есть мысли?...
test(<параметры передаваемые в функцию>)
*/
func (cat *Cats) test() float64 {
	return float64(cat.age) * cat.happines
}

func main() {
	//структуры - проще говоря, это model хранилище данных, как в классах подключающихся к БД, без методов
	bob := Cats{"Bob", 7, 0.87}
	fmt.Println(bob.name, "age is", bob.age)
	fmt.Println("function happines bob is ", bob.test())
}
