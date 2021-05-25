package threadpool

type worker struct {
	//id        int
	//done      *sync.WaitGroup
	readyPool chan chan Task
	work      chan Task
	quit      chan bool
}

//func NewWorker(id int, readyPool chan chan Work, done *sync.WaitGroup) *worker {
func NewWorker(readyPool chan chan Task) *worker {
	return &worker{
		//id:        id,
		//done:      done,
		readyPool: readyPool,
		work:      make(chan Task),
		quit:      make(chan bool),
	}
}

func (w *worker) Process(work Task) {
	work.Do()
}

func (w *worker) Start() {
	go func() {
		for {
			w.readyPool <- w.work //hey i am ready to work on new job
			select {
			case work := <-w.work: // hey i am waiting for new job
				w.Process(work) // ok i am on it
			case <-w.quit:
				return
			}
		}
	}()
}

func (w *worker) Stop() {
	//tell worker to stop after current process
	w.quit <- true
}
