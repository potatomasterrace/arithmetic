package main

import (
	"fmt"
	"runtime"
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

func mpzFromPtr(ptr unsafe.Pointer) Mpz {
	m := Mpz{
		ptr: &ptr,
	}
	runtime.SetFinalizer(m.ptr, freeMpz)
	return m
}

func (m Mpz) String() string {
	str := C.pmpz_get_str(*m.ptr, 10)
	defer free(unsafe.Pointer(str))
	goStr := C.GoString(str)
	return goStr
}

func (a Mpz) AddMpz(b Mpz) Mpz {
	ptr := C.pmpz_add(*a.ptr, *b.ptr)
	return mpzFromPtr(ptr)
}

// define clear
// https://gmplib.org/manual/Integer-Arithmetic

func CreateMpzFromString(s string, base int) Mpz {
	cStr := C.CString(s)
	defer free(unsafe.Pointer(cStr))
	cBase := C.int(base)
	ptr := C.pmpz_set_str(cStr, cBase)
	return mpzFromPtr(ptr)
}

func main() {
	str := ""
	for i, mpz := 0, CreateMpzFromString("16", 8); i < 100; i++ {
		str = mpz.String()
		mpz = mpz.AddMpz(mpz)
		fmt.Println("value", str, "ptr", *mpz.ptr)
		runtime.GC()
	}
}
