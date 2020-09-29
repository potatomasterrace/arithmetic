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

var cnt = 0

func mpzFromPtr(ptr unsafe.Pointer) Mpz {
	m := Mpz{
		ptr: &ptr,
	}
	cnt++
	localCnt := cnt
	fmt.Println("allocated : ", localCnt, m.Ptr())
	runtime.SetFinalizer(m.ptr, func(ptr *unsafe.Pointer) {
		fmt.Println("freeing : ", localCnt, *ptr)
		freeMpz(ptr)
	})
	return m
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

func runStuff() {
	str := ""
	for i, mpz := 0, CreateMpzFromString("111", 8); i < 100; i++ {
		str = mpz.String()
		for i := 0; i < 1000; i++ {
			runtime.GC()
		}
		mpz = mpz.AddMpz(mpz)
		fmt.Println("value", str, "cnt", cnt, "ptr", mpz.ptr)
	}
}

func main() {
	for i := 0; i < 10; i++ {
		runStuff()
	}
	for i := 0; i < 1000; i++ {
		runtime.GC()
	}
}
