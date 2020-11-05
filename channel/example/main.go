package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Producer struct {
	OutChan chan int
}

//вовзращает канал из которого можно только читать инкапсуляция
func (p *Producer) getOutChan() <-chan int {
	return p.OutChan
}

func (p *Producer) produce() {
	for {
		time.Sleep(3 * time.Second)
		p.OutChan <- rand.Intn(100) //если будет больше 10 элементов, канал будет подвисать здесь. ожидая освобождения памяти
	}
}

/*Channel - Chan = Канал
В каналы имеют некую очередь в которой следуют принципу FIFO First Input Fisrt Output
первый зашел, первый вышел. */
func main() {
	/*//Создаем канал, целочисленного типа, размерности 10 элементов, под котоырй будет выделено место в буфере
	intChan := make(chan int, 10)
	// передаем в созданный канал число 1
	intChan <- 1
	//Забираем элементы из канала
	fmt.Println(<-intChan)*/

	//Каналы используют для взаимодействия между потоками
	readChan := make(chan int, 10)
	prod := Producer{
		OutChan: readChan,
	}
	go prod.produce() //убрать, будет вечное ожидание пустого сообщения из канала gorutines deadline
	/*for {
		i := <-readChan
		fmt.Println("Got message from chan", i)
	}*/

	/*for i := range prod.getOutChan() {
		fmt.Println("Got message from chan", i)
	}*/

	prodChan := prod.getOutChan()
	ticker := time.NewTicker(2 * time.Second)
	for {
		// switch, неблакирующий функционал считывания из каналов
		select {
		case a := <-prodChan:
			fmt.Println("Got message from producer", a)
		case <-ticker.C:
			fmt.Println("Got message from ticker")
		}
		/*3 продюсера работающие на kafka и нужно получать данные со всех.
		Можно через select получать данные с каждого worker и реализовать для каждого канала
		свою кастомную логику*/
	}
}
