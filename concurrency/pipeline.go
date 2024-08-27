package concurrency

import "fmt"

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
