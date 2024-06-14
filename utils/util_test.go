package utils

import (
	"reflect"
	"testing"
)

type TestStruct struct {
	Name    string `json:"name" db:"Name"`
	Address string `json:"address" db:"Addr"`
	Age     uint   `json:"age" db:"Age"`
}

func TestExtractTagFromStruct(t *testing.T) {
	tests := []struct {
		name string
		s    interface{}
		tags []string
		want map[string]map[string]string
	}{
		{
			"test with json and db tag",
			&TestStruct{Name: "John", Address: "Street X", Age: 20},
			[]string{"json", "db"},
			map[string]map[string]string{
				"Name":    {"json": "name", "db": "Name"},
				"Address": {"json": "address", "db": "Addr"},
				"Age":     {"json": "age", "db": "Age"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExtractTagFromStruct(tt.s, tt.tags...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExtractTagFromStruct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetNextIndexStrings(t *testing.T) {
	cases := []struct {
		name     string
		input    []string
		expected string
		isErr    bool
		reset    bool
	}{
		{"empty", []string{}, "", false, false},
		{"single element", []string{"one"}, "one", false, false},
		{"multiple elements", []string{"one", "two", "three"}, "three", false, false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.reset {
				index.Store(0)
			}
			actual := GetNextIndexStrings(tc.input)
			if actual != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}

func TestGetNextIndex(t *testing.T) {

	tests := []struct {
		name string
		arg  int
		want int
	}{
		{"Negative Input", -5, 0},
		{"Zero Input", 0, 0},
		{"Positive Input", 7, 1},            // Assuming index start at 0
		{"Large Positive Input", 100000, 1}, // Assuming index start at 0
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index.Store(0)
			if got := GetNextIndex(tt.arg); got != tt.want {
				t.Errorf("GetNextIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByteZfill(t *testing.T) {
	type args struct {
		b []byte
		n int
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "test1",
			args: args{
				b: []byte{1, 2, 3},
				n: 5,
			},
			want: []byte{0, 0, 1, 2, 3},
		},
		{
			name: "test2",
			args: args{
				b: []byte{1, 2, 3},
				n: 2,
			},
			want: []byte{1, 2, 3},
		},
		{
			name: "test3",
			args: args{
				b: []byte{1, 2, 3},
				n: 5,
			},
			want: []byte{0, 0, 1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ZFillByte(tt.args.b, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ZFillByte() = %v, want %v", got, tt.want)
			}
		})
	}
}
