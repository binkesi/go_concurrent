package concurrent_models

import "time"

func GenerateNum() chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()
	return ch
}

func PrimeFilter(in <-chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				out <- i
			}
		}
	}()
	return out
}

var Done = false
var A string

func Setup() {
	A = "Hello world"
	time.Sleep(time.Second)
	Done = true
}
