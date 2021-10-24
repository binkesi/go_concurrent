package concurrent_models

import (
	"context"
	"fmt"
	"testing"
	"time"
)

const DurationTime = 1 * time.Millisecond

func TestContextTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), DurationTime)
	defer cancel()
	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}

func TestTimeAfter(t *testing.T) {
	c := make(chan int)
	go func() {
		c <- 1
	}()
	select {
	case m := <-c:
		fmt.Printf("%v\n", m)
	case <-time.After(5 * time.Second):
		fmt.Println("time out")
	}
}
