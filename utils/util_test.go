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

func TestHexToInt(t *testing.T) {
	type args struct {
		hex string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"0x1", args{"0x1"}, "1"},
		{"0xff", args{"0xff"}, "255"},
		{"0xffff", args{"0xffff"}, "65535"},
		{"0xffffff", args{"0xffffff"}, "16777215"},
		{"0xffffffff", args{"0xffffffff"}, "4294967295"},
		{"0xffffffffff", args{"0xffffffffff"}, "1099511627775"},
		{"0xffffffffffff", args{"0xffffffffffff"}, "281474976710655"},
		{"0xffffffffffffff", args{"0xffffffffffffff"}, "72057594037927935"},
		{"0xffffffffffffffff", args{"0xffffffffffffffff"}, "18446744073709551615"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HexToString(tt.args.hex); got != tt.want {
				t.Errorf("HexToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
