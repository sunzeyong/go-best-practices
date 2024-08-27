package concurrency

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapReduce(t *testing.T) {
	stream := asStream(nil)

	mapFn := func(v interface{}) interface{} {
		return v.(int) * 10
	}
	reduceFn := func(a, b interface{}) interface{} {
		return a.(int) + b.(int)
	}

	sum := reduce(mapChan(stream, mapFn), reduceFn)
	assert.Equal(t, 150, sum)
}
