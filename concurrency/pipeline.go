package concurrency

import (
	"fmt"
	"sync"
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

// 优化，流水线上某个操作提高处理人手，即扇入扇出模型
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
