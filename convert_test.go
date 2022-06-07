package tokenizer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBytesToString(t *testing.T) {
	tests := []struct {
		name string
		args []byte
		want string
	}{
		{
			"test bytes to string conversion english, numbers & symbols",
			[]byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz01234567890!@#$%^&*()-=_+{}[]:;'\",./<>?`~"),
			"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz01234567890!@#$%^&*()-=_+{}[]:;'\",./<>?`~",
		},
		{
			"test bytes to string conversion russian, numbers & symbols",
			[]byte("АБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯабвгдеёжзийклмнопрстуфхцчшщъыьэюя01234567890!@#$%^&*()-=_+{}[]:;'\",./<>?`~"),
			"АБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯабвгдеёжзийклмнопрстуфхцчшщъыьэюя01234567890!@#$%^&*()-=_+{}[]:;'\",./<>?`~",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, bytesToString(tt.args), tt.want)
		})
	}
}

func TestStringToBytes(t *testing.T) {
	tests := []struct {
		name string
		args string
		want []byte
	}{
		{
			"test string to bytes conversion english, numbers & symbols",
			"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz01234567890!@#$%^&*()-=_+{}[]:;'\",./<>?`~",
			[]byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz01234567890!@#$%^&*()-=_+{}[]:;'\",./<>?`~"),
		},
		{
			"test string to bytes conversion russian, numbers & symbols",
			"АБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯабвгдеёжзийклмнопрстуфхцчшщъыьэюя01234567890!@#$%^&*()-=_+{}[]:;'\",./<>?`~",
			[]byte("АБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯабвгдеёжзийклмнопрстуфхцчшщъыьэюя01234567890!@#$%^&*()-=_+{}[]:;'\",./<>?`~"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, stringToBytes(tt.args), tt.want)
		})
	}
}

func BenchmarkBytesToString(b *testing.B) {
	b.ReportAllocs()
	testRunes := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz01234567890!@#$%^&*()-=_+{}[]:;'\",./<>?`~")

	for i := 0; i < b.N; i++ {
		_ = bytesToString(testRunes)
	}
}

func BenchmarkStringToBytes(b *testing.B) {
	b.ReportAllocs()
	testRunes := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz01234567890!@#$%^&*()-=_+{}[]:;'\",./<>?`~"

	for i := 0; i < b.N; i++ {
		_ = stringToBytes(testRunes)
	}
}
