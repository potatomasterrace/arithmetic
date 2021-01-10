package mpz

import (
	"runtime"
	"unsafe"
)

// #cgo LDFLAGS: -lgmp -L${SRCDIR}/memory
// #cgo CFLAGS: -I${SRCDIR}/memory -Wno-incompatible-pointer-types
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
	defer runtime.SetFinalizer(m.ptr, func(ptr *unsafe.Pointer) {
		go C.clear_mpz(*ptr)
	})
	return m
}
