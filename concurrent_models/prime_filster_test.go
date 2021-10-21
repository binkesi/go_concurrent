package concurrentmodels

import (
	"fmt"
	"testing"
)

func TestPrimeFilter(t *testing.T) {
	ch := GenerateNum()
	for i := 0; i < 100; i++ {
		prime := <-ch
		fmt.Printf("%v: %v\n", i+1, prime)
		ch = PrimeFilter(ch, prime)
	}
}
