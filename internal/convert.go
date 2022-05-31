package internal

import (
	"reflect"
	"unsafe"
)

// BytesToString is zero-allocation converter from byte-slice to string
func BytesToString(bytes []byte) string {
	return *(*string)(unsafe.Pointer(&bytes))
}

// StringToBytes is zero-allocation converter from string to byte-slice
func StringToBytes(str string) []byte {
	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&str))
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: stringHeader.Data,
		Len:  stringHeader.Len,
		Cap:  stringHeader.Len,
	}))
}
