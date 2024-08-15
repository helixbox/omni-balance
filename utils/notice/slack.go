package notice

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"omni-balance/utils"

	"go.uber.org/zap/zapcore"
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

func Level2SlackColor(level zapcore.Level) string {
	switch level {
	case zapcore.ErrorLevel:
		return "danger"
	case zapcore.WarnLevel:
		return "warning"
	default:
		return "good"
	}
}

func (s *Slack) Send(ctx context.Context, title string, content string, level zapcore.Level, fields Fields) error {
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
	for k, v := range fields {
		if utils.InArrayFold(k, []string{"title", "message"}) {
			continue
		}
		if v == "" || k == "" {
			continue
		}
		b.Attachments[0].Fields = append(b.Attachments[0].Fields, Field{
			Title: k,
			Value: v,
			Short: true,
		})
	}
	var bodyBuf = bytes.NewBuffer(nil)
	if err := json.NewEncoder(bodyBuf).Encode(b); err != nil {
		return err
	}
	_ = utils.Request(ctx, http.MethodPost, s.Webhook, bodyBuf, nil)
	return nil
}
