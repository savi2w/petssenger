package worker

import (
	"github.com/vmihailenco/taskq/v2"
	"github.com/vmihailenco/taskq/v2/redisq"
	"github.com/weslenng/petssenger/services/pricing/config"
	"github.com/weslenng/petssenger/services/pricing/models"
	"github.com/weslenng/petssenger/services/pricing/redis"
)

var (
	// QueueFactory is a asynchronous task/job queue with Redis
	QueueFactory = redisq.NewFactory()
	// MainQueue is a main queue of QueueFactory
	MainQueue = QueueFactory.RegisterQueue(&taskq.QueueOptions{
		Name:  "pricing-worker",
		Redis: redis.Client,
	})
	// DecreaseDynamicFees must decrease dynamic fees on all requests after 5 minutes
	DecreaseDynamicFees = taskq.RegisterTask(&taskq.TaskOptions{
		Name:       "decrease-dynamic-fees",
		Handler:    models.DecreaseDynamicFees,
		RetryLimit: config.Default.WorkerRetryLimit,
		MinBackoff: config.Default.WorkerRetryDelay,
	})
)
