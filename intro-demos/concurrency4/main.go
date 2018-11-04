package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// START worker OMIT
func print(ctx context.Context, worker, count int, ch chan time.Duration) {
	start := time.Now()
	defer func() {
		ch <- time.Since(start)
	}()

	for count > 0 {
		fmt.Printf("[worker %3d] %3d bottles of mate on the wall\n", worker, count)
		timeout := time.After(time.Duration(count) * 50 * time.Millisecond)
		count--
		select {
		case <-ctx.Done():
			fmt.Printf("[worker %3d] cancelled\n", worker)
			return
		case <-timeout:
		}
	}
}

// END worker OMIT

// START main OMIT
func main() {
	ch := make(chan time.Duration)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	for i := 0; i < 15; i++ {
		go print(ctx, i, rand.Intn(20), ch)
	}

	for i := 0; i < 15; i++ {
		d := <-ch
		fmt.Printf("print() finished after %v\n", d)
	}
}

// END main OMIT
