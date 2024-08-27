package concurrency

import (
	"reflect"
	"time"
)

// 当等待一组goroutine全部执行完可以用waitgroup
// 若等待一组goroutine其中一个执行完就算完成，就是orDone模式 属于channel信号通知一种使用场景

// 返回orDone
func or(ins ...<-chan struct{}) <-chan struct{} {
	if len(ins) == 0 {
		return nil
	}

	orDone := make(chan struct{})
	// 监听ins 若有事件产生，则close(orDone)
	go func() {
		defer close(orDone)

		cases := make([]reflect.SelectCase, 0, len(ins))
		for _, in := range ins {
			cases = append(cases, reflect.SelectCase{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(in),
			})
		}
		reflect.Select(cases)
	}()

	return orDone
}

func sig(duration time.Duration) <-chan struct{} {
	c := make(chan struct{})

	go func() {
		time.Sleep(duration)
		close(c)
	}()
	return c
}
