package worker

import (
	"time"

	"github.com/vmihailenco/taskq/v2"
	"github.com/vmihailenco/taskq/v2/redisq"
	"github.com/weslenng/petssenger/services/pricing/config"
	"github.com/weslenng/petssenger/services/pricing/models"
)

var (
	// QueueFactory is a asynchronous task/job queue with Redis
	QueueFactory = redisq.NewFactory()
	// MainQueue is a main queue of QueueFactory
	MainQueue = QueueFactory.RegisterQueue(&taskq.QueueOptions{
		Name:  "pricing-worker",
		Redis: config.PricingRedisClient(),
	})
	// DecreaseDynamicFees must decrease dynamic fees on all requests after 5 minutes
	DecreaseDynamicFees = taskq.RegisterTask(&taskq.TaskOptions{
		Name:       "decrease-dynamic-fees",
		Handler:    models.DecreaseDynamicFees,
		RetryLimit: 3,
		MinBackoff: 3 * time.Second,
	})
)
