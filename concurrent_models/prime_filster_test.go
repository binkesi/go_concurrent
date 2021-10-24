package concurrent_models

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPrimeFilter(t *testing.T) {
	ch := GenerateNum()
	for i := 0; i < 100; i++ {
		prime := <-ch
		fmt.Printf("%v: %v\n", i+1, prime)
		ch = PrimeFilter(ch, prime)
	}
}

func TestRandomNum(t *testing.T) {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			select {
			case ch <- 0:
			case ch <- 1:
			}
		}
		close(ch)
	}()
	for in := range ch {
		fmt.Println(in)
	}
}

func Worker(wg *sync.WaitGroup, cancel chan interface{}) {
	defer wg.Done()
	for {
		select {
		default:
			time.Sleep(time.Second)
			fmt.Println("Hello world!")
		case <-cancel:
			return
		}
	}
}

func TestCancel(t *testing.T) {
	ch := make(chan interface{})
	wg := new(sync.WaitGroup)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go Worker(wg, ch)
	}
	time.Sleep(3 * time.Second)
	close(ch)
	wg.Wait()
}
