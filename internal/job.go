package internal

import (
	"context"
	"fmt"
	"github.com/axieinfinity/bridge-v2/internal/types"
	"github.com/axieinfinity/bridge-v2/internal/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"time"
)

type Job struct {
	ID         int32
	Type       int
	Message    []interface{}
	RetryCount int32
	NextTry    int32
	MaxTry     int32
	BackOff    int32
	Listener   types.IListener
}

func (job Job) Hash() common.Hash {
	return common.BytesToHash([]byte(fmt.Sprintf("j-%d-%d-%d", job.ID, job.RetryCount, job.NextTry)))
}

type Worker struct {
	ctx context.Context

	utilWrapper utils.IUtils

	id int

	// queue is passed from subscriber is used to add workerChan to queue
	queue chan chan types.IJob

	// mainChain is controller's jobChan which is used to push job back to controller
	mainChan chan<- types.IJob

	// workerChan is used to receive and process job
	workerChan chan types.IJob

	failedChan  chan<- types.IJob
	successChan chan<- types.IJob

	closeChan chan struct{}
	listeners map[string]types.IListener
}

func NewWorker(ctx context.Context, id int, mainChan, failedChan, successChan chan<- types.IJob, queue chan chan types.IJob, size int, listeners map[string]types.IListener) *Worker {
	return &Worker{
		ctx:         ctx,
		id:          id,
		workerChan:  make(chan types.IJob, size),
		mainChan:    mainChan,
		failedChan:  failedChan,
		successChan: successChan,
		queue:       queue,
		listeners:   listeners,
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

func (w *Worker) processJob(job types.IJob) {
	var (
		err error
		val []byte
	)
	val, err = job.Process()
	if err != nil {
		log.Error("[Worker] failed while processing job", "id", job.GetID(), "err", err)
		goto ERROR
	}
	if job.GetType() == types.ListenHandler {
		job.GetListener().SendCallbackJobs(w.listeners, job.GetSubscriptionName(), job.GetTransaction(), val, w.mainChan)
	}
	w.successChan <- job
	return
ERROR:
	if job.GetRetryCount()+1 > job.GetMaxTry() {
		log.Info("[Worker][processJob] job reaches its maxTry", "jobTransaction", job.GetTransaction().GetHash().Hex())
		w.failedChan <- job
		return
	}
	job.IncreaseRetryCount()
	job.UpdateNextTry(time.Now().Unix() + int64(job.GetRetryCount()*job.GetBackOff()))
	// push the job back to mainChan
	w.mainChan <- job
	log.Info("[Worker][processJob] job failed, added back to jobChan", "id", job.GetID(), "retryCount", job.GetRetryCount(), "nextTry", job.GetNextTry())
}
