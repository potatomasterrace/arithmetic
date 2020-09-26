package main

import "unsafe"

// #cgo LDFLAGS: -lgmp
// #include "mpz.h"
import "C"

// Conversion Functions
// char *pmpz_get_str(const unsafe_mpz n, const int base);
func (m Mpz) GetString(base C.int) string {
	str := C.pmpz_get_str(*m.ptr, base)
	defer free(unsafe.Pointer(str))
	goStr := C.GoString(str)
	return goStr
}

func (m Mpz) String() string {
	return m.GetString(10)
}

// double pmpz_get_d(const unsafe_mpz op);
func (m Mpz) GetDouble() C.double {
	return C.pmpz_get_d(*m.ptr)
}

// double pmpz_get_d_2exp(signed long int *exp, const unsafe_mpz op);
func (m Mpz) GetD2Exp() (val C.double, exp C.long) {
	val = C.pmpz_get_d_2exp(&exp, *m.ptr)
	return val, exp
}

// signed long int pmpz_get_si(const unsafe_mpz op);
func (m Mpz) GetSignedInt() C.long {
	return C.pmpz_get_si(*m.ptr)
}

// unsigned long int pmpz_get_ui(const unsafe_mpz op);
func (m Mpz) GetUnsignedInt() C.ulong {
	return C.pmpz_get_ui(*m.ptr)
}
