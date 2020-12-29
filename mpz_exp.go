package main

// #cgo LDFLAGS: -lgmp
// #include "mpz.h"
import "C"

// Exponentiation Functions
// unsafe_mpz pmpz_pow_ui(const unsafe_mpz base, unsigned long int exp);
func (m Mpz) MpzPowUi(bitcnt C.ulong) Mpz {
	ptr := C.pmpz_fdiv_r_2exp(m.Ptr(), bitcnt)
	return mpzFromPtr(ptr)
}

// unsafe_mpz pmpz_powm(const unsafe_mpz base, const unsafe_mpz exp, const unsafe_mpz mod);
func (m Mpz) MpzPowM(exp Mpz, mod Mpz) Mpz {
	ptr := C.pmpz_powm(m.Ptr(), exp.Ptr(), mod.Ptr())
	return mpzFromPtr(ptr)
}

// unsafe_mpz pmpz_powm_sec(const unsafe_mpz base, const unsafe_mpz exp, const unsafe_mpz mod);
func (m Mpz) MpzPowMSec(exp Mpz, mod Mpz) Mpz {
	ptr := C.pmpz_powm_sec(m.Ptr(), exp.Ptr(), mod.Ptr())
	return mpzFromPtr(ptr)
}

// unsafe_mpz pmpz_powm_ui(const unsafe_mpz base, unsigned long int exp, const unsafe_mpz mod);
func (m Mpz) MpzPowMUi(exp C.ulong, mod Mpz) Mpz {
	ptr := C.pmpz_powm_ui(m.Ptr(), exp, mod.Ptr())
	return mpzFromPtr(ptr)
}
