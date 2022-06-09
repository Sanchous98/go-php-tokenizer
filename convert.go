package tokenizer

import (
	"reflect"
	"unsafe"
)

// bytesToString is zero-allocation converter from byte-slice to string
func bytesToString(bytes []byte) string {
	return *(*string)(unsafe.Pointer(&bytes))
}

// stringToBytes is zero-allocation converter from string to byte-slice
func stringToBytes(str string) []byte {
	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&str))
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: stringHeader.Data,
		Len:  stringHeader.Len,
		Cap:  stringHeader.Len,
	}))
}

func noescape[T any](value *T) *T {
	v := uintptr(unsafe.Pointer(value))
	return (*T)(unsafe.Pointer(v ^ 0))
}
