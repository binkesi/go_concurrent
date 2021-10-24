package concurrent_models

import (
	"fmt"
	"testing"
)

var done = false
var a string

func TestGoroutine(t *testing.T) {
	go Setup()
	for !done {
	}
	fmt.Println(a)
}
