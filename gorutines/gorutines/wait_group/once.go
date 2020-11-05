package wait_group

import (
	"fmt"
	"sync"
)

func main() {
	once := sync.Once{}
	for i := 0; i < 10; i++ {
		once.Do(Print())
	}
}

func Print() {
	fmt.Println("Hello World!")
}
