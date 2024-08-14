package pool

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	// init pool
	pool := NewPool(2)
	pool.Run()

	// start request
	reqNo := 10
	go func() {
		for i := 0; i < reqNo; i++ {
			jobItem := &JobExample{
				Num: i,
			}
			pool.TaskQ <- jobItem
		}
	}()

	for {
		fmt.Println("the number of goroutine: ", runtime.NumGoroutine())
		time.Sleep(2 * time.Second)
	}

}
