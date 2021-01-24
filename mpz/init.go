package mpz

import (
	"unsafe"
)

// #cgo LDFLAGS: -lgmp
// #include "mpz.h"
import "C"

// Initialisation and assignement

func MpzInit() Mpz {
	ptr := C.pmpz_init()
	m := mpzFromPtr(ptr)
	return m
}

var cnt = 0

// unsafe_mpz pmpz_set_str(const char *s, const int base)Mpz{ }
func MpzFromString(s string, base int) Mpz {
	cStr := C.CString(s)
	defer C.free(unsafe.Pointer(cStr))
	cBase := C.int(base)
	ptr := C.pmpz_set_str(cStr, cBase)
	return mpzFromPtr(ptr)
}

// unsafe_mpz pmpz_set_double(const double op)Mpz{ }
func MpzFromDouble(op float64) Mpz {
	ptr := C.pmpz_set_double(C.double(op))
	return mpzFromPtr(ptr)
}

// unsafe_mpz pmpz_set_mpz(const unsafe_mpz op)Mpz{ }
func MpzFromMpz(op Mpz) Mpz {
	ptr := C.pmpz_set_mpz(op.Ptr())
	return mpzFromPtr(ptr)
}

// unsafe_mpz pmpz_set_si(const signed long int op)Mpz{ }
func MpzFromSi(op C.long) Mpz {
	ptr := C.pmpz_set_si(op)
	return mpzFromPtr(ptr)
}

// unsafe_mpz pmpz_set_ui(const unsigned long int op)Mpz{ }
func MpzFromUi(op C.ulong) Mpz {
	ptr := C.pmpz_set_ui(op)
	return mpzFromPtr(ptr)
}

// unsafe_mpz pmpz_ui_pow_ui(unsigned long int base, unsigned long int exp);
func (m Mpz) MpzFromUiPowUi(base C.ulong, exp C.ulong) Mpz {
	ptr := C.pmpz_ui_pow_ui(base, exp)
	return mpzFromPtr(ptr)
}

// unsafe_mpz pmpz_2fac_ui(unsigned long int n);
func Mpz2FacUi(n C.ulong) Mpz {
	ptr := C.pmpz_2fac_ui(n)
	return mpzFromPtr(ptr)
}

// unsafe_mpz pmpz_fac_ui(unsigned long int n);
func MpzFacUi(n C.ulong) Mpz {
	ptr := C.pmpz_fac_ui(n)
	return mpzFromPtr(ptr)
}

// unsafe_mpz pmpz_fib_ui(unsigned long int n);
func MpzFibUi(n C.ulong) Mpz {
	ptr := C.pmpz_fib_ui(n)
	return mpzFromPtr(ptr)
}

// unsafe_mpz pmpz_lucnum_ui(unsigned long int n);
func MpzLucNumUi(n C.ulong) Mpz {
	ptr := C.pmpz_lucnum_ui(n)
	return mpzFromPtr(ptr)
}

// unsafe_mpz pmpz_mfac_uiui(unsigned long int n, unsigned long int m);
func MpzMFacUiUi(n C.ulong, m C.ulong) Mpz {
	ptr := C.pmpz_mfac_uiui(n, m)
	return mpzFromPtr(ptr)
}

// unsafe_mpz pmpz_primorial_ui(unsigned long int n);
func MpzPrimordialUi(n C.ulong) Mpz {
	ptr := C.pmpz_primorial_ui(n)
	return mpzFromPtr(ptr)
}

// TODO After mpq
// unsafe_mpz pmpz_set_quotient(const mpq_t op)Mpz{ }
//func MpzFromQuotient(const mpq_t op) Mpz {
//	ptr := pmpz_set_quotient(*op.ptr)
//	return mpzFromPtr(ptr)
//}
// unsafe_mpz pmpz_set_float(const mpf_t op)Mpz{ }
