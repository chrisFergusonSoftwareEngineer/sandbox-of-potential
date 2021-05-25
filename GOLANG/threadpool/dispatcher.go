package threadpool

import "fmt"

type Pool struct {
	internalQueue chan Task
	readyPool     chan chan Task //boss says hey i have a new job at my desk workers who available can get it in this way he does not have to ask current status of workers
	workers       []*worker
	quit          chan bool
}

var workerPool *Pool

func GetWorkerPool() *Pool {
	if nil == workerPool {
		workerPool = newWorkerPool(10, 1000)
	}

	return workerPool
}

func newWorkerPool(maxWorkers int, jobQueueCapacity int) *Pool {
	if jobQueueCapacity <= 0 {
		jobQueueCapacity = 10000
	}

	readyPool := make(chan chan Task, maxWorkers)
	workers := make([]*worker, maxWorkers, maxWorkers)

	// create workers
	for i := 0; i < maxWorkers; i++ {
		workers[i] = NewWorker(readyPool)
	}

	return &Pool{
		internalQueue: make(chan Task, jobQueueCapacity),
		readyPool:     readyPool,
		workers:       workers,
		quit:          make(chan bool),
	}
}

func (q *Pool) Start() {
	//tell workers to get ready
	for i := 0; i < len(q.workers); i++ {
		q.workers[i].Start()
	}
	// open factory
	go q.dispatch()
}

func (q *Pool) Stop() {
	q.quit <- true
}

func (q *Pool) dispatch() {
	//open factory gate
	for {
		select {
		case job := <-q.internalQueue:
			workerXChannel := <-q.readyPool //free worker x founded
			workerXChannel <- job           // here is your job worker x
		case <-q.quit:
			// free all workers
			for i := 0; i < len(q.workers); i++ {
				q.workers[i].Stop()
			}
			return
		}
	}
}

/*Tries to enqueue but fails if queue is full*/
func (q *Pool) Enqueue(job Task) bool {
	select {
	case q.internalQueue <- job:
		return true
	default:
		fmt.Println("########## ERROR - failed to enque task!!!!!")
		//panic(nil)
		return false
	}
}
