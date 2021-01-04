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
func freeMpz(ptr *unsafe.Pointer) {
	C.clear_mpz(*ptr)
}
