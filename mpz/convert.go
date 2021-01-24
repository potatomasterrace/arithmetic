package mpz

import (
	"unsafe"
)

// #include "mpz.h"
// #include <malloc.h>
import "C"

// Conversion Functions
// char *pmpz_get_str(const unsafe_mpz n, const int base);
func (m Mpz) GetString(base int) (string, error) {
	mPtr, err := m.Ptr()
	if err != nil {
		return "", err
	}
	str := C.pmpz_get_str(mPtr, C.int(base))
	defer C.free(unsafe.Pointer(str))
	goStr := C.GoString(str)
	return goStr, nil
}

func (m Mpz) GetStringOrPanic(base int) string {
	str, err := m.GetString(base)
	if err != nil {
		panic(err)
	}
	return str
}

func (m Mpz) String() (string, error) {
	return m.GetString(10)
}
func (m Mpz) StringOrPanic() string {
	str, err := m.String()
	if err != nil {
		panic(err)
	}
	return str
}

// double pmpz_get_d(const unsafe_mpz op);
func (m Mpz) GetDouble() (float64, error) {
	mPtr, err := m.Ptr()
	if err != nil {
		return 0, err
	}
	ret := C.pmpz_get_d(mPtr)
	return float64(ret), nil
}

func (m Mpz) GetDoubleOrPanic() float64 {
	ret, err := m.GetDouble()
	if err != nil {
		panic(err)
	}
	return ret
}

// double pmpz_get_d_2exp(signed long int *exp, const unsafe_mpz op);
func (m Mpz) GetD2Exp() (v float64, e int, err error) {
	mPtr, err := m.Ptr()
	if err != nil {
		return 0, 0, err
	}
	expC := C.long(0)
	val := C.pmpz_get_d_2exp(&expC, mPtr)
	return float64(val), int(expC), nil
}

func (m Mpz) GetD2ExpOrPanic() (v float64, e int) {
	v, e, err := m.GetD2Exp()
	if err != nil {
		panic(err)
	}
	return v, e
}

// signed long int pmpz_get_si(const unsafe_mpz op);
func (m Mpz) GetSignedInt() (int, error) {
	mPtr, err := m.Ptr()
	if err != nil {
		return 0, err
	}
	ret := C.pmpz_get_si(mPtr)
	return int(ret), nil
}

func (m Mpz) GetSignedIntOrPanic() int {
	ret, err := m.GetSignedInt()
	if err != nil {
		panic(err)
	}
	return ret
}

// unsigned long int pmpz_get_ui(const unsafe_mpz op);
func (m Mpz) GetUnsignedInt() (uint, error) {
	mPtr, err := m.Ptr()
	if err != nil {
		return 0, err
	}
	ret := C.pmpz_get_ui(mPtr)
	return uint(ret), nil
}

func (m Mpz) GetUnsignedIntOrPanic() uint {
	ret, err := m.GetUnsignedInt()
	if err != nil {
		panic(err)
	}
	return ret
}
