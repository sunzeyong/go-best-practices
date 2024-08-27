package concurrency

import (
	"fmt"
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

func TestDispatchNum(t *testing.T) {
	dispatchNum()
}

func TestDispatchNum01(t *testing.T) {
	dispatchNum01()
}
