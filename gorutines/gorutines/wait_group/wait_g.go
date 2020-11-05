package wait_group

import (
	"fmt"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)
	for i := 0; i < 10; i++ {
		go Print("Hello", wg)
	}
	wg.Wait()
}

func Print(word string, wg *sync.WaitGroup) {
	wg.Add(2)
	defer wg.Done()
	fmt.Println(word)
}
