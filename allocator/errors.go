package allocator

import "errors"

var (
	ErrInvalidSize  = errors.New("malloc :size must be greater than 0")
	ErrOutOfMemory  = errors.New("malloc:out of memory")
	ErrInvalidFree  = errors.New("free:pointer was not returned by malloc")
	ErrDoubleFree   = errors.New("free:double free detected")
	ErrHeapTooSmall = errors.New("malloc:heap size is too small")
)
