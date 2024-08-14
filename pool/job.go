package pool

import (
	"fmt"
	"time"
)

type Job interface {
	Do()
}

type JobExample struct {
	Num int
}

func (j *JobExample) Do() {
	fmt.Printf("this task no: %v\n", j.Num)
	time.Sleep(time.Second)
}
