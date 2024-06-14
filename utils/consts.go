package utils

import "sync/atomic"

var (
	// 是否初始化完成
	isInit atomic.Bool
)

func IsFinishedInit() bool {
	return isInit.Load()
}

func FinishInit() {
	isInit.Store(true)
}
