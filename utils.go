package flexbuffers

import "C"
import (
	"bytes"
	"fmt"
	"reflect"
	"unsafe"
)

func stringToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}

func cstringBytesToString(s []byte) (string, error) {
	l := bytes.IndexByte(s, 0)
	if l < 0 {
		return "", fmt.Errorf("null not found")
	}
	ss := s[0:l]
	return *(*string)(unsafe.Pointer(&ss)), nil
}

func bytesToString(s []byte) string {
	return *(*string)(unsafe.Pointer(&s))
}

func unsafeBufferString(buf []byte, offset, size int) string {
	var sh reflect.StringHeader
	sh.Len = size
	sh.Data = uintptr(unsafe.Pointer(&buf[offset]))
	return *(*string)(unsafe.Pointer(&sh))
}

func readCStringBytes(buf []byte, offset int) []byte {
	size := bytes.IndexByte(buf[offset:], 0)
	if size < 0 {
		panic("no null terminator")
	}
	return buf[offset : offset+size]
}

func unsafeReadCString(buf []byte, offset int) (string, error) {
	if offset < 0 || len(buf) <= offset {
		return "", ErrOutOfRange
	}
	size := bytes.IndexByte(buf[offset:], 0)
	if size < 0 {
		return "", fmt.Errorf("null not found")
	}
	return unsafeBufferString(buf, offset, size), nil
}
