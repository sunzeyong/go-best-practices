package concurrency

import (
	"fmt"
	"os"
	"os/signal"
	"reflect"
	"sync"
	"syscall"
	"time"
)

// Pipeline 模式也称为流水线模式，模拟的就是现实世界中的流水线生产。
// 从技术上看，每一道工序的输出，就是下一道工序的输入。
// 在工序之间传递的东西就是数据，这种模式称为流水线模式，而传递的数据称为数据流。

// 采购
func buy(n int) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for i := 0; i < n; i++ {
			out <- fmt.Sprintf("配件i:%v", i)
		}
	}()
	return out
}

// 组装
func build(in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for item := range in {
			out <- fmt.Sprintf("组装(%v)", item)
		}
	}()
	return out
}

// 打包
func pack(in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for item := range in {
			out <- fmt.Sprintf("打包(%v)", item)
		}
	}()
	return out
}

// 优化，流水线上某个操作提高处理人手，即扇出扇入模型
// 由于channel是并发安全的，所以扇出只需要将上一个channel发送给多个下一个处理工序
// 而扇入则需要添加merge函数，将多个channel合并为一个channel
func merge(ins ...<-chan string) <-chan string {
	out := make(chan string)
	var wg sync.WaitGroup

	send := func(in <-chan string) {
		defer wg.Done()
		for item := range in {
			out <- item
		}
	}

	wg.Add(len(ins))
	for _, ch := range ins {
		go send(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

// 扇入模型的另一种写法
func fanIn(ins ...<-chan string) <-chan string {
	out := make(chan string)

	go func() {
		defer close(out)
		// 组装select cases
		cases := make([]reflect.SelectCase, 0, len(out))
		for _, item := range ins {
			cases = append(cases, reflect.SelectCase{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(item),
			})
		}

		// 监听cases
		for len(cases) > 0 {
			i, v, ok := reflect.Select(cases)
			if !ok {
				cases = append(cases[:i], cases[i+1:]...)
				continue
			}
			out <- v.String()
		}
	}()

	return out
}

// 更普遍的一种方式，任意的goroutine数量之间传递数据

// 开一百个协程打印数字和当前的协程No
// 协程1打印以1结尾的数，协程2打印以2结尾的数，。。。，协程100打印100结尾的数
// 打印出1-1000 并且按照顺序打印

// sig用于终止打印
func printWorker(i int, pre, next chan int, sig chan struct{}) {
	for v := range pre {
		fmt.Printf("goroutine: %v, num: %v\n", i, v)

		if v == 1000 {
			close(sig)
		} else {
			v++
			next <- v
		}
	}
}

func dispatchNum() {
	chs := make([]chan int, 100)
	for i := range chs {
		chs[i] = make(chan int)
	}

	sig := make(chan struct{})

	for i := 1; i <= 100; i++ {
		go printWorker(i, chs[i-1], chs[(i)%100], sig)
	}

	chs[0] <- 1

	<-sig
	fmt.Println("game is over")

}

// 若上述问题简化

// 有四个 goroutine，编号为 1、2、3、4。每秒钟会有一个 goroutine 打印出它自己的编号，
// 要求你编写一个程序，让输出的编号总是按照 1、2、3、4、1、2、3、4、……的顺序打印出来。

func printWorker01(i int, pre, next chan struct{}) {
	for {
		token := <-pre
		fmt.Printf("goroutine: %v\n", i+1)
		time.Sleep(time.Second)
		next <- token
	}
}

func dispatchNum01() {
	chs := make([]chan struct{}, 4)
	for i := range chs {
		chs[i] = make(chan struct{})
	}

	token := struct{}{}
	for i := 0; i < 4; i++ {
		go printWorker01(i, chs[i], chs[(i+1)%4])
	}
	chs[0] <- token

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Printf("goroutine is over by %v\n", <-sig)
}
