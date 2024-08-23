package concurrency

import (
	"fmt"
)

// 开一百个协程打印数字和当前的协程No
// 协程1打印以1结尾的数，协程2打印以2结尾的数，。。。，协程100打印100结尾的数
// 打印出1-1000 并且按照顺序打印

func PrintWorker(i int, pre, next chan int, sig chan struct{}) {
	for {
		num := <-pre
		fmt.Printf("groutine: %v, %v\n", i, num)

		if num == 1000 {
			close(sig)
		} else {
			num++
			next <- num
		}
	}
}

func PrintDispatch() {
	ch := make([]chan int, 100)
	for i := range ch {
		ch[i] = make(chan int)
	}

	sig := make(chan struct{})

	for i := 0; i < 100; i++ {
		num := i + 1
		go PrintWorker(num, ch[i], ch[(i+1)%100], sig)
	}
	ch[0] <- 1

	<-sig
}
