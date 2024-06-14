package notice

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"omni-balance/utils"
)

type Slack struct {
	Webhook string `json:"webhook"`
	Channel string `json:"channel"`
}

type body struct {
	Channel     string       `json:"channel"`
	Username    string       `json:"username"`
	UnfurlLinks bool         `json:"unfurl_links"`
	Attachments []Attachment `json:"attachments"`
}

type Attachment struct {
	Fallback   string  `json:"fallback"`
	Color      string  `json:"color"`
	AuthorName string  `json:"author_name"`
	Fields     []Field `json:"fields"`
}
type Field struct {
	Title string `json:"title"`
	Value string `json:"value"`
	Short bool   `json:"short"`
}

func Level2SlackColor(level logrus.Level) string {
	switch level {
	case logrus.ErrorLevel:
		return "danger"
	case logrus.WarnLevel:
		return "warning"
	default:
		return "good"
	}
}

func (s *Slack) Send(ctx context.Context, title string, content string, levels ...logrus.Level) error {
	var level = logrus.DebugLevel
	if len(levels) > 0 {
		level = levels[0]
	}
	b := &body{
		Channel:     s.Channel,
		Username:    "omni-balance",
		UnfurlLinks: false,
		Attachments: []Attachment{
			{
				Fallback: fmt.Sprintf("%s\n%s", title, content),
				Color:    Level2SlackColor(level),
				Fields: []Field{
					{
						Title: "title",
						Value: title,
						Short: true,
					},
					{
						Title: "message",
						Value: content,
						Short: true,
					},
				},
			},
		},
	}
	var bodyBuf = bytes.NewBuffer(nil)
	if err := json.NewEncoder(bodyBuf).Encode(b); err != nil {
		return err
	}
	_ = utils.Request(ctx, http.MethodPost, s.Webhook, bodyBuf, nil, "Content-Type", "application/json")
	return nil
}
