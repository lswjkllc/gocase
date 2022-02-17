package main

import (
	"fmt"
	"time"
)

var MAX_WORKER int = 3
var MAX_QUEUE int = 100

// func init() {
// 	MAX_WORKER, _ = strconv.Atoi(os.Getenv("MAX_WORKER"))
// 	MAX_QUEUE, _ = strconv.Atoi(os.Getenv("MAX_QUEUE"))
// }

type Payload struct {
	// [redacted]
	// storageFolder string
}

func (p *Payload) UploadToS3() error {
	time.Sleep(200 * time.Millisecond)
	fmt.Println("Upload success...")
	return nil
	// the storageFolder method ensures that there are no name collision in
	// case we get same timestamp in the key name
	// storage_path := fmt.Sprintf("%v/%v", p.storageFolder, time.Now().UnixNano())

	// bucket := S3Bucket

	// b := new(bytes.Buffer)
	// encodeErr := json.NewEncoder(b).Encode(payload)
	// if encodeErr != nil {
	// 	return encodeErr
	// }

	// // Everything we post to the S3 bucket should be marked 'private'
	// var acl = s3.Private
	// var contentType = "application/octet-stream"

	// return bucket.PutReader(storage_path, b, int64(b.Len()), contentType, acl, s3.Options{})
}

// Job represents the job to be run
type Job struct {
	Payload Payload
}

// Worker represents the worker that executes the job
type Worker struct {
	WorkerPool chan chan Job
	JobChannel chan Job
	quit       chan bool
}

func NewWorker(workerPool chan chan Job) Worker {
	return Worker{
		WorkerPool: workerPool,
		JobChannel: make(chan Job),
		quit:       make(chan bool)}
}

// Start method starts the run loop for the worker, listening for a quit channel in
// case we need to stop it
func (w Worker) Start() {
	go func() {
		for {
			// register the current worker into the worker queue.
			// 生产 chan Job
			w.WorkerPool <- w.JobChannel

			select {
			case job := <-w.JobChannel: // 消费 Job
				// we have received a work request.
				if err := job.Payload.UploadToS3(); err != nil {
					fmt.Printf("Error uploading to S3: %s", err.Error())
				}

			case <-w.quit:
				// we have received a signal to stop
				fmt.Println("Received stop signal")
				return
			}
		}
	}()
}

// Stop signals the worker to stop listening for work requests.
func (w Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}

///
type Dispatcher struct {
	maxWorkers int
	// A pool of workers channels that are registered with the dispatcher
	WorkerPool chan chan Job
}

func NewDispatcher(maxWorkers int) *Dispatcher {
	pool := make(chan chan Job, maxWorkers)
	return &Dispatcher{WorkerPool: pool, maxWorkers: maxWorkers}
}

func (d *Dispatcher) Run() {
	// starting n number of workers
	for i := 0; i < d.maxWorkers; i++ {
		worker := NewWorker(d.WorkerPool)
		worker.Start()
	}

	// go d.dispatch()
	go d.dispatch()
}

func (d *Dispatcher) dispatch() {
	for job := range JobQueue {
		// a job request has been received
		go func(job Job) {
			// try to obtain a worker job channel that is available.
			// this will block until a worker is idle
			// 消费 chan Job
			jobChannel := <-d.WorkerPool

			// dispatch the job to the worker job channel
			// 生产 Job
			jobChannel <- job
		}(job)
	}
	// for {
	// 	select {
	// 	case job := <-JobQueue:
	// 		// a job request has been received
	// 		go func(job Job) {
	// 			// try to obtain a worker job channel that is available.
	// 			// this will block until a worker is idle
	// 			jobChannel := <-d.WorkerPool

	// 			// dispatch the job to the worker job channel
	// 			jobChannel <- job
	// 		}(job)
	// 		// default:
	// 		// 	fmt.Println("waiting...")
	// 	}
	// }
}

// A buffered channel that we can send work requests on.
var JobQueue chan Job = make(chan Job, MAX_QUEUE)

func main() {
	// 初始化 dispatcher
	dispatcher := NewDispatcher(MAX_WORKER)
	// 开始运行
	dispatcher.Run()

	for {
		JobQueue <- Job{Payload{}}
		time.Sleep(2 * time.Second)
	}
}
