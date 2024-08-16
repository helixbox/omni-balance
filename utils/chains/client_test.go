package chains

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/spf13/cast"
)

func Test_needTry(t *testing.T) {
	tests := []struct {
		err  error
		want bool
	}{
		{nil, false}, // 测试 nil 错误
		{errors.New("connection refused"), true},
		{errors.New("no such host"), true},
		{errors.New("i/o timeout"), true},
		{errors.New("some other error"), false},
	}

	for _, tt := range tests {
		t.Run(cast.ToString(tt.err), func(t *testing.T) {
			if got := needTry(tt.err); got != tt.want {
				t.Errorf("needTry() = %v, want %v", got, tt.want)
			}
		})
	}
}
