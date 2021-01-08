package mpz

import (
	"unsafe"
)

// #cgo LDFLAGS: -lgmp
// #include "mpz.h"
import "C"

type Mpz struct {
	ptr *unsafe.Pointer
}

func (a Mpz) Ptr() unsafe.Pointer {
	return *a.ptr
}

func mpzFromPtr(ptr unsafe.Pointer) Mpz {
	m := Mpz{
		ptr: &ptr,
	}
	defer RegisterPtrReference(m.ptr, func(ptr *unsafe.Pointer) {
		C.clear_mpz(*ptr)
	})
	return m
}
