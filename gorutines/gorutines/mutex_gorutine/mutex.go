package mutex_gorutine

import "sync"

type MyObj struct {
	MapMutex *sync.Mutex
	Map      map[string]string
}

func main() {
	obj := MyObj{
		Map:      make(map[string]string, 0),
		MapMutex: new(sync.Mutex),
	}
	writeToMap(&obj)
}

func writeToMap(obj *MyObj) {
	obj.MapMutex.Lock()
	defer obj.MapMutex.Unlock()
	obj.Map["hello"] = "123"
}
