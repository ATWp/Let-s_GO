package fib

var (
	cache = map[int]int
)

func init(){
	cache = make(map[int]int)
}

func Fib(n int) int {
	if n <= 1 {
		return n
	}

	if r, ok := cache[n]; ok{
		return r
	}
	
	sum := Fib(n-1) + Fib(n-2)
	cache[n] = sum

	return sum
}
