package locks

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"omni-balance/utils"
	"strings"
	"sync"
	"time"
)

var (
	nameMutexMap = make(map[string]*sync.Mutex)
	mutex        sync.Mutex
)

// LockKey generate lock key, if values is empty, generate a random string
func LockKey(values ...any) string {
	if len(values) == 0 {
		values = append(values, uuid.New().String())
	}
	var format []string
	for range values {
		format = append(format, "%s")
	}
	return strings.ToLower(fmt.Sprintf(strings.Join(format, "_"), values...))
}

func LockWithKey(ctx context.Context, key string, noWait ...bool) bool {
	log := utils.GetLogFromCtx(ctx)
	tryLock := func() bool {
		mutex.Lock()
		defer mutex.Unlock()
		locker, ok := nameMutexMap[key]
		if !ok {
			nameMutexMap[key] = &sync.Mutex{}
			nameMutexMap[key].Lock()
			return true
		}
		return locker.TryLock()
	}
	ok := tryLock()
	if len(noWait) != 0 && noWait[0] && !ok {
		return false
	}
	if ok {
		return true
	}
	var t = time.NewTicker(time.Second)
	defer t.Stop()
	for {
		select {
		case <-ctx.Done():
			return false
		case <-t.C:
			if tryLock() {
				return true
			}
			if log != nil {
				log.Debugf("%s is locked, waiting...", key)
			}
		}
	}
}

func UnlockWithKey(_ context.Context, key string) {
	mutex.Lock()
	defer mutex.Unlock()
	if _, ok := nameMutexMap[key]; !ok {
		nameMutexMap[key] = &sync.Mutex{}
		return
	}
	nameMutexMap[key].Unlock()
}
