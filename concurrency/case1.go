package concurrency

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// 有四个 goroutine，编号为 1、2、3、4。每秒钟会有一个 goroutine 打印出它自己的编号，
// 要求你编写一个程序，让输出的编号总是按照 1、2、3、4、1、2、3、4、……的顺序打印出来。

func worker(i int, pre, next chan int) {
	for v := range pre {
		n := v % 4
		if i-1 == n {
			fmt.Println(i)
		} else {
			next <- v
		}
	}
}

func master() {
	ch := make([]chan int, 5)
	for i := range ch {
		ch[i] = make(chan int)
	}

	for i := 1; i < 5; i++ {
		go worker(i, ch[i-1], ch[i])
	}

	go func() {
		ticker := time.NewTicker(time.Second)
		total := 0
		for {
			<-ticker.C
			ch[0] <- total
			total++
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Printf("goroutine is over by %v\n", <-sig)
}
