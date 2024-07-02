package market

import "omni-balance/utils/bot"

var (
	taskQueue = make(chan Task, 100)
)

type Task struct {
	Id          string
	ProcessType bot.ProcessType
}

func PushTask(task Task) {
	taskQueue <- task
}
