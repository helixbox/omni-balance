package bot

import (
	"omni-balance/internal/daemons"
	"time"
)

func init() {
	daemons.RegisterIntervalTask(daemons.Task{
		Name:            "bot",
		Description:     "Check the balance based on the specific bot and create an order",
		TaskFunc:        Run,
		DefaultInterval: time.Minute * 3,
	})
}
