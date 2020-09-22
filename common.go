package main

import (
	"unsafe"
)

// #include <malloc.h>
import "C"

func free(ptr unsafe.Pointer) {
	go C.free(ptr)
}
