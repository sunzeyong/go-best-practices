package oberver

import (
	"fmt"
	"testing"
	"time"
)

func O1(msg1, msg2 string) {
	fmt.Printf("O1, msg1:%s, msg2:%s\n", msg1, msg2)
}

func O2(msg1, msg2 string) {
	fmt.Printf("O2, msg1:%s, msg:%s\n", msg1, msg2)
}

func TestEventBus(t *testing.T) {
	bus := NewAsyncEventBus()
	bus.Subscribe("topic:1", O1)
	bus.Subscribe("topic:2", O2)

	bus.Public("topic:1", "test1", "test2")
	bus.Public("topic:1", "testA", "testB")

	time.Sleep(2*time.Second)
}
