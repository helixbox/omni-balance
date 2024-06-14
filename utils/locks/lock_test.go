package locks

import (
	"context"
	"testing"
)

func TestLockKey(t *testing.T) {
	type args struct {
		values []any
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{
				values: []any{"test1", "test2"},
			},
			want: "test1_test2",
		},
		{
			name: "test2",
			args: args{
				values: []any{"test1", "test2", "test3"},
			},
			want: "test1_test2_test3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LockKey(tt.args.values...); got != tt.want {
				t.Errorf("LockKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLockWithKey(t *testing.T) {
	type args struct {
		ctx    context.Context
		key    string
		noWait bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{
				ctx: context.Background(),
				key: "test1",
			},
			want: true,
		},
		{
			name: "noWaitLock",
			args: args{
				ctx:    context.Background(),
				key:    "test1",
				noWait: true,
			},
			want: false,
		},
		{
			name: "noWaitLock",
			args: args{
				ctx:    context.Background(),
				key:    "test1",
				noWait: true,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LockWithKey(tt.args.ctx, tt.args.key, tt.args.noWait); got != tt.want {
				t.Errorf("LockWithKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
