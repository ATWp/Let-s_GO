package main

import "fmt"

func main() {
	c := make(chan string)

	go ping_pong(c)
	for i := 0; i < 5; i++ {
		fmt.Println(<-c, ",", <-c)
	}

	rangeChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		rangeChan <- i
	}
	close(rangeChan)
	fmt.Println("Sum:", sum(rangeChan))
}

func ping_pong(c chan<- string) {
	for {
		c <- fmt.Sprintf("Ping")
		c <- fmt.Sprintf("Pong")
	}
}

func sum(input <-chan int) (res int) {
	for r := range input {
		res += r
	}
	return res
}
