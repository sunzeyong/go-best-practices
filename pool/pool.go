package pool

type Pool struct {
	Len     int
	TaskQ   chan Job
	WorkerQ chan chan Job
}

func NewPool(len int) *Pool {
	return &Pool{
		Len:     len,
		TaskQ:   make(chan Job),
		WorkerQ: make(chan chan Job, len),
	}
}

func (p *Pool) Run() {
	// 初始化pool中worker
	for i := 0; i < p.Len; i++ {
		w := NewWorker()
		w.Run(p.WorkerQ)
	}

	// pool 工作流程
	go func() {
		for {
			// 1. 有任务么
			job := <-p.TaskQ
			// 2. 有空闲worker么
			wTaskChan := <-p.WorkerQ
			// 3. 任务分配给worker
			wTaskChan <- job
		}
	}()
}
