package main

import "fmt"

//запуск через новый тип
type myString string

func (*myString) process(i int) {
	fmt.Println("process struct", i)
}

//Отдельная функция процесса
func process(i int) {
	fmt.Println("process:", i)
}

func main() {
	fmt.Println("Start goroutine")
	//запуск простой функции
	go process(1)

	//анонимная функция lambda
	go func() {
		fmt.Println("Anon func")
	}()

	//запуск в цикле
	for i := 0; i < 10; i++ {
		go process(i)
	}

	//запуск метода типа
	myStr := new(myString)
	go myStr.process(1)

	//Дожидаемся завершения выполнения и ловим первые ошибки запуска
	// handler для остановки - функция main завершается и убивает все горутины
	fmt.Scanln()
}
