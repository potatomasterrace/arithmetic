package mpz

import (
	"fmt"
	"runtime"
	"unsafe"
)

// #cgo LDFLAGS: -lgmp -L${SRCDIR}/memory
// #cgo CFLAGS: -I${SRCDIR}/memory -Wno-incompatible-pointer-types
// #include "mpz.h"
import "C"

// Stub function used for testing.
var clearMpz func(ptr *unsafe.Pointer) = func(ptr *unsafe.Pointer) {
	go C.clear_mpz(*ptr)
}

type Mpz struct {
	ptr *unsafe.Pointer
}

func (a Mpz) Ptr() unsafe.Pointer {
	return *a.ptr
}

func mpzFromPtr(ptr unsafe.Pointer) (Mpz, error) {
	if ptr != nil {
		m := Mpz{
			ptr: &ptr,
		}
		defer runtime.SetFinalizer(m.ptr, clearMpz)
		return m, nil
	}
	return Mpz{}, fmt.Errorf("cannot initialize value")
}
