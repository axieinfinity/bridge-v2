package metrics

import (
	"context"

	"github.com/axieinfinity/bridge-v2/adapters/prometheus"
)

const (
	EthereumProcessedBlockMetric string = "ethereum/processedBlock"

	PreparingFailedJobMetric  string = "jobs/prepare/failed"
	PreparingSuccessJobMetric string = "jobs/prepare/success"
	ProcessingJobMetric       string = "jobs/processing"
	ProcessedSuccessJobMetric string = "jobs/success"
	ProcessedFailedJobMetric  string = "jobs/failed"

	PendingTaskMetric       string = "Ronin/tasks/pending"
	ProcessingTaskMetric    string = "Ronin/tasks/processing"
	SuccessTaskMetric       string = "Ronin/tasks/success"
	FailedTaskMetric        string = "Ronin/tasks/failed"
	WithdrawalTaskMetric    string = "Ronin/tasks/withdrawal"
	DepositTaskMetric       string = "Ronin/tasks/deposit"
	AckWithdrawalTaskMetric string = "Ronin/tasks/acknowledgeWithdrawal"
)

var (
	Pusher *prometheus.PusherAdapter
)

func RunPusher(ctx context.Context) {
	Pusher = prometheus.NewPusher()
	Pusher.AddCounter(PreparingFailedJobMetric, "count number of preparing jobs failed to added to database")
	Pusher.AddCounter(PreparingSuccessJobMetric, "count number of preparing jobs added to database successfully and switch to new ")
	Pusher.AddCounter(ProcessingJobMetric, "count number of processing jobs in jobChan")
	Pusher.AddCounter(ProcessedSuccessJobMetric, "count number of processed jobs successfully")
	Pusher.AddCounter(ProcessedFailedJobMetric, "count number of failed jobs")
	Pusher.AddCounter(PendingTaskMetric, "count number of pending tasks in queue")
	Pusher.AddCounter(ProcessingJobMetric, "count number of processing tasks")
	Pusher.AddCounter(ProcessedSuccessJobMetric, "count number of success tasks")
	Pusher.AddCounter(ProcessedFailedJobMetric, "count number failed tasks")
	Pusher.AddCounter(WithdrawalTaskMetric, "count number of ronin’s withdrawal tasks occurred")
	Pusher.AddCounter(DepositTaskMetric, "count number of ronin’s deposit tasks occurred")
	Pusher.AddCounter(AckWithdrawalTaskMetric, "count number of ronin acknowledge withdrawal tasks occurred")

	go Pusher.Start(ctx)
}
