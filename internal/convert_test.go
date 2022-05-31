package internal

import (
	"reflect"
	"testing"
)

func TestBytesToString(t *testing.T) {
	tests := []struct {
		name string
		args []byte
		want string
	}{
		{"test bytes to string conversion", []byte("Test"), "Test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BytesToString(tt.args); got != tt.want {
				t.Errorf("BytesToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRunesToString(t *testing.T) {
	tests := []struct {
		name string
		args []rune
		want string
	}{
		{"test runes to string conversion", []rune("Test"), "Test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RunesToString(tt.args); got != tt.want {
				t.Errorf("RunesToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringToBytes(t *testing.T) {
	tests := []struct {
		name string
		args string
		want []byte
	}{
		{"test bytes to string conversion", "Test", []byte("Test")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringToBytes(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringToBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}
