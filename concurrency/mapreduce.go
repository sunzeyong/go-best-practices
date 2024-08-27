package concurrency

// map-reduce 是一种处理数据的方式，最早是由 Google 公司研究提出的一种面向大规模数据处理的并行计算模型和方法
// map-reduce 分为两个步骤:
// 第一步是映射（map），处理队列中的数据，
// 第二步是规约（reduce），把列表中的每一个元素按照一定的处理方式处理成结果，放入到结果队列中。

// map 处理队列中数据，返回channel，里面会包含处理后的数据
func mapChan(in <-chan interface{}, f func(interface{}) interface{}) <-chan interface{} {
	out := make(chan interface{})

	if in == nil {
		close(out)
		return out
	}

	go func() {
		defer close(out)
		for v := range in {
			out <- f(v)
		}
	}()

	return out
}

// 处理mapchan中数据
func reduce(in <-chan interface{}, f func(interface{}, interface{}) interface{}) interface{} {
	if in == nil {
		return nil
	}

	out := <-in
	for v := range in {
		out = f(out, v)
	}
	return out
}

func asStream(done chan struct{}) <-chan interface{} {
	out := make(chan interface{})

	values := []int{1, 2, 3, 4, 5}
	go func() {
		defer close(out)

		for _, v := range values {
			select {
			case <-done:
				return
			case out <- v:
			}
		}
	}()

	return out
}
