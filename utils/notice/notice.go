package notice

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/patrickmn/go-cache"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"omni-balance/utils"
	"omni-balance/utils/constant"
	"sync"
	"time"
)

var (
	msgInterval = time.Hour
	msg         = cache.New(time.Hour, time.Minute)
	m           sync.Mutex
)

type Type string

const (
	SlackNotice Type = "slack"
)

var (
	notice Notice
)

type Fields map[string]string

type Notice interface {
	Send(ctx context.Context, title string, content string, level logrus.Level, fields Fields) error
}

func SetMsgInterval(interval time.Duration) {
	if interval.Seconds() < time.Hour.Seconds() {
		logrus.Warnf("msg interval %s is too short, set to 1 hour", interval)
		msgInterval = time.Hour
		return
	}
	msgInterval = interval
}

func WithFields(ctx context.Context, fields Fields) context.Context {
	return context.WithValue(ctx, constant.NoticeFieldsKeyInCtx, fields)
}

func Init(noticeType Type, conf map[string]interface{}, interval time.Duration) error {
	if notice != nil {
		return nil
	}
	confData, _ := json.Marshal(conf)
	switch noticeType {
	case SlackNotice:
		s := &Slack{}
		if err := json.Unmarshal(confData, s); err != nil {
			return errors.Wrap(err, "unmarshal slack config")
		}
		notice = s
	default:
		if noticeType != "" {
			return errors.Errorf("notice type %s not support", noticeType)
		}
	}
	SetMsgInterval(interval)
	return nil
}

func Send(ctx context.Context, title string, content string, levels ...logrus.Level) error {
	m.Lock()
	defer m.Unlock()
	if notice == nil {
		return nil
	}

	key := utils.Md5(fmt.Sprintf("%s:%s", title, content))
	if _, ok := msg.Get(key); ok {
		logrus.Debugf("notice %s:%s already send, 1 hour later will send again", title, content)
		return nil
	}
	if len(levels) == 0 {
		levels = append(levels, logrus.InfoLevel)
	}
	var fields Fields
	value := ctx.Value(constant.NoticeFieldsKeyInCtx)
	if v, ok := value.(Fields); ok {
		fields = v
	}
	if err := notice.Send(ctx, title, content, levels[0], fields); err != nil {
		return err
	}
	msg.Set(key, struct{}{}, msgInterval)
	return nil
}
