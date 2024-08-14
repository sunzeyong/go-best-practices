package pool

type Worker struct {
	TaskChan chan Job
	Quit     chan struct{}
}

func NewWorker() Worker {
	return Worker{
		TaskChan: make(chan Job),
		Quit:     make(chan struct{}),
	}
}

func (w *Worker) Run(pq chan chan Job) {
	go func() {
		for {
			// 1. 初次跑起来，将自己的task channel注册到pool的task queue中
			// 2. 每次任务执行完 再次注册上去
			pq <- w.TaskChan
			// 注册完之后就只需要关心自己的task channel是否有任务进来，没有就阻塞
			select {
			case job := <-w.TaskChan:
				job.Do()
			case <-w.Quit:
				return
			}
		}
	}()
}
