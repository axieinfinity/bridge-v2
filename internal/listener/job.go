package listener

import (
	"context"
	"fmt"
	"github.com/axieinfinity/bridge-v2/internal/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"time"
)

const (
	listenHandler = iota
	callbackHandler
)

type Job struct {
	ID         int32
	Type       int
	Message    []interface{}
	RetryCount int32
	NextTry    int32
	MaxTry     int32
	BackOff    int32
	Listener   IListener
}

func (job Job) Hash() common.Hash {
	return common.BytesToHash([]byte(fmt.Sprintf("j-%d-%d-%d", job.ID, job.RetryCount, job.NextTry)))
}

type Worker struct {
	ctx context.Context

	utilWrapper utils.IUtils

	id int

	// queue is passed from subscriber is used to add workerChan to queue
	queue chan chan IJob

	// mainChain is controller's jobChan which is used to push job back to controller
	mainChan chan<- IJob

	// workerChan is used to receive and process job
	workerChan chan IJob

	closeChan chan struct{}
	listeners map[string]IListener
}

func NewWorker(ctx context.Context, id int, mainChan chan<- IJob, queue chan chan IJob, size int, listeners map[string]IListener) *Worker {
	return &Worker{
		ctx:        ctx,
		id:         id,
		workerChan: make(chan IJob, size),
		mainChan:   mainChan,
		queue:      queue,
		listeners:  listeners,
	}
}

func (w *Worker) String() string {
	return fmt.Sprintf("{ id: %d, currentSize: %d }", w.id, len(w.workerChan))
}

func (w *Worker) start() {
	for {
		// push worker chan into queue
		w.queue <- w.workerChan
		select {
		case job := <-w.workerChan:
			log.Info("processing job", "id", job.GetID(), "nextTry", job.GetNextTry(), "retryCount", job.GetRetryCount(), "type", job.GetType())
			if job.GetNextTry() == 0 || job.GetNextTry() <= time.Now().Unix() {
				w.processJob(job)
				continue
			}
			// push the job back to mainChan
			w.mainChan <- job
		case <-w.ctx.Done():
			close(w.workerChan)
			return
		}
	}
}

func (w *Worker) processJob(job IJob) {
	var (
		err error
		val interface{}
	)
	switch job.GetType() {
	case listenHandler:
		val, err = job.Process()
		if err != nil {
			goto ERROR
		}
		job.GetListener().SendCallbackJobs(w.listeners, job.GetSubscriptionName(), job.GetTransaction(), val, w.mainChan)
	case callbackHandler:
		_, err = job.Process()
		if err != nil {
			goto ERROR
		}
	}
	return
ERROR:
	if job.GetRetryCount()+1 > job.GetMaxTry() {
		log.Info("[Worker][processJob] job reaches its maxTry", "jobTransaction", job.GetTransaction().GetHash().Hex())
		return
	}
	job.IncreaseRetryCount()
	job.UpdateNextTry(time.Now().Unix() + int64(job.GetRetryCount()*job.GetBackOff()))
	// push the job back to mainChan
	w.mainChan <- job
	log.Info("[Worker][processJob] job failed, added back to jobChan", "id", job.GetID(), "retryCount", job.GetRetryCount(), "nextTry", job.GetNextTry())
}
