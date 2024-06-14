package notice

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/patrickmn/go-cache"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"omni-balance/utils"
	"time"
)

var (
	msg = cache.New(time.Hour, time.Minute)
)

type Type string

const (
	SlackNotice Type = "slack"
)

var (
	notice Notice
)

type Notice interface {
	Send(ctx context.Context, title string, content string, levels ...logrus.Level) error
}

func Init(noticeType Type, conf map[string]interface{}) error {
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
	return nil
}

func Send(ctx context.Context, title string, content string, levels ...logrus.Level) error {
	if notice == nil {
		return nil
	}

	key := utils.Md5(fmt.Sprintf("%s:%s", title, content))
	if _, ok := msg.Get(key); ok {
		logrus.Debugf("notice %s:%s already send, 1 hour later will send again", title, content)
		return nil
	}
	if err := notice.Send(ctx, title, content, levels...); err != nil {
		return err
	}
	msg.SetDefault(key, struct{}{})
	return nil
}
