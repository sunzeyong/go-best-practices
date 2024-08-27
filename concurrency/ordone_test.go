package concurrency

import (
	"fmt"
	"testing"
	"time"
)

func TestOrDone(t *testing.T) {
	st := time.Now()

	<-or(
		sig(1*time.Second),
		sig(2*time.Second),
		sig(3*time.Second))

	fmt.Printf("timeout: %v", time.Since(st))
}
