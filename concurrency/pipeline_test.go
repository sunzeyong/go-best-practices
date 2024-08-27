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
