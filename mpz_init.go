package main

// #cgo LDFLAGS: -lgmp
// #include "mpz.h"
import "C"
import (
	"unsafe"
)

// Initialisation and assignement

func MpzInit() Mpz {
	ptr := C.pmpz_init()
	return mpzFromPtr(ptr)
}

// unsafe_mpz pmpz_set_str(const char *s, const int base)Mpz{ }
func MpzFromString(s string, base int) Mpz {
	cStr := C.CString(s)
	defer free(unsafe.Pointer(cStr))
	cBase := C.int(base)
	ptr := C.pmpz_set_str(cStr, cBase)
	return mpzFromPtr(ptr)
}

// unsafe_mpz pmpz_set_double(const double op)Mpz{ }
func MpzFromDouble(op C.double) Mpz {
	ptr := C.pmpz_set_double(op)
	return mpzFromPtr(ptr)
}

// unsafe_mpz pmpz_set_mpz(const unsafe_mpz op)Mpz{ }
func MpzFromMpz(op Mpz) Mpz {
	ptr := C.pmpz_set_mpz(op.Ptr())
	return mpzFromPtr(ptr)
}

// TODO After mpq
// unsafe_mpz pmpz_set_quotient(const mpq_t op)Mpz{ }
//func MpzFromQuotient(const mpq_t op) Mpz {
//	ptr := pmpz_set_quotient(*op.ptr)
//	return mpzFromPtr(ptr)
//}
// unsafe_mpz pmpz_set_float(const mpf_t op)Mpz{ }

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
