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

func (m Mpz) MpzMod(bitcnt C.ulong) Mpz {
	ptr := C.pmpz_fdiv_r_2exp(m.Ptr(), bitcnt)
	return mpzFromPtr(ptr)
}

// unsafe_mpz pmpz_powm_sec(const unsafe_mpz base, const unsafe_mpz exp, const unsafe_mpz mod);

func (m Mpz) MpzMod(bitcnt C.ulong) Mpz {
	ptr := C.pmpz_fdiv_r_2exp(m.Ptr(), bitcnt)
	return mpzFromPtr(ptr)
}

// unsafe_mpz pmpz_powm_ui(const unsafe_mpz base, unsigned long int exp, const unsafe_mpz mod);

func (m Mpz) MpzMod(bitcnt C.ulong) Mpz {
	ptr := C.pmpz_fdiv_r_2exp(m.Ptr(), bitcnt)
	return mpzFromPtr(ptr)
}

// unsafe_mpz pmpz_ui_pow_ui(unsigned long int base, unsigned long int exp);

func (m Mpz) MpzMod(bitcnt C.ulong) Mpz {
	ptr := C.pmpz_fdiv_r_2exp(m.Ptr(), bitcnt)
	return mpzFromPtr(ptr)
}
