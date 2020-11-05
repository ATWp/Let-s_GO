package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
)

var ( //кэш наших номеров чисел последовательности и ответов, чисел на определенном месте
	cache map[int]int
)

//неизменяемое значение порта на котором весит сервер
const PORT string = ":8080"

//инициализируем матрицу cache
func init() {
	cache = make(map[int]int)
}

//Основная функция для запуска сервера
func main() {
	//все роутеры - ссылки
	http.HandleFunc("/fib", handleFib)
	//для удобства, выводим ссылку на локальный сервер, который поднимаем
	fmt.Println("Server start on: http://127.0.0.1" + PORT + "/fib")
	//запускаем сервер, слушая заданный порт
	http.ListenAndServe(PORT, nil)
}

//Функция обработки request запросов по адресу "/fib"
func handleFib(w http.ResponseWriter, r *http.Request) {
	//создаем целочисленную переменную в которую складываем переданное через ссылку значение /fib?num=<значение>
	num := r.FormValue("num")
	//преобразуем полученное в виде строки число в тип string с помощью сторонней библиотеки strconv
	n, err := strconv.Atoi(num)
	//ловим ошибку преобразования
	if err != nil {
		//возвращаем ошибку преобразования и код ошибки BadRequest
		http.Error(w, err.Error(), http.StatusBadRequest)
		// выходим из функции
		return
	}

	//передаем обратно полученное число Фибоначчи из функции fib сразу преобразовывая его в строку библиотекой strconv
	io.WriteString(w, strconv.Itoa(fib(n)))
}

//функция фибоначи
func fib(n int) int {
	if n <= 1 {
		return n
	}

	if r, ok := cache[n]; ok {
		return r
	}

	sum := fib(n-1) + fib(n-2)
	cache[n] = sum
	return sum
}
