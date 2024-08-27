package concurrency

import (
	"fmt"
	"sync"
	"testing"
)

func TestPipeline(t *testing.T) {
	buyChan := buy(10)
	buildChan := build(buyChan)
	packChan := pack(buildChan)

	for item := range packChan {
		fmt.Println(item)
	}
}

func TestMerge(t *testing.T) {
	buyChan := buy(10)

	build01 := build(buyChan)
	build02 := build(buyChan)
	mergeBuild := merge(build01, build02)

	packChan := pack(mergeBuild)

	for item := range packChan {
		fmt.Println(item)
	}
}

func TestFanIn(t *testing.T) {
	buyChan := buy(10)

	build01 := build(buyChan)
	build02 := build(buyChan)
	mergeBuild := fanIn(build01, build02)

	packChan := pack(mergeBuild)

	for item := range packChan {
		fmt.Println(item)
	}
}

func TestFanOut(t *testing.T) {
	in := make(chan interface{})

	outs := make([]chan interface{}, 5)
	for i := range outs {
		outs[i] = make(chan interface{})
	}

	fanOut(in, outs, false)

	go func() {
		defer close(in)
		for i := range 3 {
			in <- i
		}
	}()

	var wg sync.WaitGroup

	for i := range outs {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			for range outs[i] {
				fmt.Printf("goroutine: %v recv again\n", i)
			}
		}()
	}
	wg.Wait()
}

func TestDispatchNum(t *testing.T) {
	dispatchNum()
}

func TestDispatchNum01(t *testing.T) {
	dispatchNum01()
}
