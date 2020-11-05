package main

import "fmt"

func main() {
	webSites := make(map[string]float64)

	webSites["test"] = 0.8
	webSites["Yandex"] = 0.99
	fmt.Println(webSites["test"])
}
