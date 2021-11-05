package worker

import (
	"github.com/vmihailenco/taskq/v2"
	"github.com/vmihailenco/taskq/v2/redisq"
	"github.com/omppye-lab/petssenger/services/pricing/config"
	"github.com/omppye-lab/petssenger/services/pricing/models"
	"github.com/omppye-lab/petssenger/services/pricing/redis"
)

var (
	QueueFactory = redisq.NewFactory()
	MainQueue    = QueueFactory.RegisterQueue(&taskq.QueueOptions{
		Name:  "pricing-worker",
		Redis: redis.Client,
	})
	DecreaseDynamicFees = taskq.RegisterTask(&taskq.TaskOptions{
		Name:       "decrease-dynamic-fees",
		Handler:    models.DecreaseDynamicFees,
		RetryLimit: config.Default.WorkerRetryLimit,
		MinBackoff: config.Default.WorkerRetryDelay,
	})
)
