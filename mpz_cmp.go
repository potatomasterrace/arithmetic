package main

// #cgo LDFLAGS: -lgmp
// #include "mpz.h"
import "C"

// Comparison Functions
// int pmpz_cmp(const unsafe_mpz op1, const unsafe_mpz op2);
func (m Mpz) CmpMpz(mpz Mpz) C.int {
	return C.pmpz_cmp(m.Ptr(), mpz.Ptr())
}

// int pmpz_cmp_d(const unsafe_mpz op1, double op2);
func (m Mpz) CmpDMpz(val C.double) C.int {
	return C.pmpz_cmp_d(m.Ptr(), val)
}

// int pmpz_cmp_si(const unsafe_mpz op1, signed long int op2);
func (m Mpz) CmpSiMpz(val C.long) C.int {
	return C.pmpz_cmp_si(m.Ptr(), val)
}

// int pmpz_cmp_ui(const unsafe_mpz op1, unsigned long int op2);
func (m Mpz) CmpUiMpz(val C.ulong) C.int {
	return C.pmpz_cmp_ui(m.Ptr(), val)
}

// int pmpz_cmpabs(const unsafe_mpz op1, const unsafe_mpz op2);
func (m Mpz) CmpAbsMpz(mpz Mpz) C.int {
	return C.pmpz_cmpabs(m.Ptr(), mpz.Ptr())
}

// int pmpz_cmpabs_d(const unsafe_mpz op1, double op2);
func (m Mpz) CmpAbsDMpz(val C.double) C.int {
	return C.pmpz_cmpabs_d(m.Ptr(), val)
}

// int pmpz_cmpabs_ui(const unsafe_mpz op1, unsigned long int op2);
func (m Mpz) CmpAbsUiMpz(val C.ulong) C.int {
	return C.pmpz_cmpabs_ui(m.Ptr(), val)
}

// int pmpz_sgn(const unsafe_mpz op);
func (m Mpz) SignMpz() C.int {
	return C.pmpz_sgn(m.Ptr())
}
