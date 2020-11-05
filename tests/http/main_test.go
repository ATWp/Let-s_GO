package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandleFib(t *testing.T) {
	/*создаем структуру для тестирования
	name имя теста
	num  запрос, который передаем. В данном случае номер числа из последовательности Фибоначчи
	want []byte ответ, который мы должны были получить, после срабатыванияалгоритма. Он будет возвращен в виде байт.
	*/
	testCases := []struct {
		name string
		num  int
		want []byte
	}{ // тесты
		{
			name: "zero",
			num:  0,
			want: []byte("0"),
		},
		{
			name: "one",
			num:  1,
			want: []byte("1"),
		},
		{
			name: "two",
			num:  2,
			want: []byte("1"),
		},
		{
			name: "three",
			num:  3,
			want: []byte("2"),
		},
	}

	//вовзращает интерфейс запроса
	handler := http.HandlerFunc(handleFib)
	//проходимся в цикле по тестовым данным
	for _, tc := range testCases {
		//подключаем будущее тестирование
		t.Run(tc.name, func(t *testing.T) {
			//псевдо интерфейс
			rec := httptest.NewRecorder()
			//псевдо запрос
			req, _ := http.NewRequest("GET", fmt.Sprintf("/fib?num=%d", tc.num), nil)
			//делаем запрос через handler к серверу
			handler.ServeHTTP(rec, req)
			//смотрим то, что вернул нам сервер на запрос сравнивая. Проверка с помощью библиотеки testify
			assert.Equal(t, tc.want, rec.Body.Bytes())
		})
	}
}
