package mpz

import (
	"unsafe"
)

// #include "mpz.h"
// #include <malloc.h>
import "C"

// Conversion Functions
// char *pmpz_get_str(const unsafe_mpz n, const int base);
func (m Mpz) GetString(base int) string {
	str := C.pmpz_get_str(m.Ptr(), C.int(base))
	defer C.free(unsafe.Pointer(str))
	goStr := C.GoString(str)
	return goStr
}

func (m Mpz) String() string {
	return m.GetString(10)
}

// double pmpz_get_d(const unsafe_mpz op);
func (m Mpz) GetDouble() float64 {
	ret := C.pmpz_get_d(m.Ptr())
	return float64(ret)
}

// double pmpz_get_d_2exp(signed long int *exp, const unsafe_mpz op);
func (m Mpz) GetD2Exp() (v float64, e int) {
	expC := C.long(0)
	val := C.pmpz_get_d_2exp(&expC, m.Ptr())
	return float64(val), int(expC)
}

// signed long int pmpz_get_si(const unsafe_mpz op);
func (m Mpz) GetSignedInt() int {
	ret := C.pmpz_get_si(m.Ptr())
	return int(ret)
}

// unsigned long int pmpz_get_ui(const unsafe_mpz op);
func (m Mpz) GetUnsignedInt() uint {
	ret := C.pmpz_get_ui(m.Ptr())
	return uint(ret)
}
