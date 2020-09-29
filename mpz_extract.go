package main

// #cgo LDFLAGS: -lgmp
// #include "mpz.h"
import "C"

// int pmpz_perfect_power_p(const unsafe_mpz op);
func (m Mpz) MpzPerfectPowerP() C.int {
	return C.pmpz_perfect_power_p(m.Ptr())
}

// int pmpz_perfect_square_p(const unsafe_mpz op);
func (m Mpz) MpzPerfectSquareP() C.int {
	return C.pmpz_perfect_square_p(m.Ptr())
}

// int pmpz_root(unsafe_mpz rop, const unsafe_mpz op, unsigned long int n);
func (m Mpz) MpzRoot() C.int {
	return C.pmpz_perfect_square_p(m.Ptr())
}

// unsafe_mpz pmpz_sqrt(const unsafe_mpz op);
func (m Mpz) MpzSqrt() Mpz {
	ptr := C.pmpz_sqrt(m.Ptr())
	return mpzFromPtr(ptr)
}

// void pmpz_rootrem(unsafe_mpz root, unsafe_mpz rem, const unsafe_mpz u, unsigned long int n);
func (m Mpz) MpzRootRem(n C.ulong) (root Mpz, rem Mpz) {
	root, rem = MpzInit(), MpzInit()
	C.pmpz_rootrem(root.Ptr(), rem.Ptr(), m.Ptr(), n)
	return root, rem
}

// void pmpz_sqrtrem(unsafe_mpz rop1, unsafe_mpz rop2, const unsafe_mpz op);
func (m Mpz) MpzSqrtRem() (root Mpz, rem Mpz) {
	rop1, rop2 = MpzInit(), MpzInit()
	ptr := C.pmpz_sqrtrem(rop1.Ptr(), rop2.Ptr(), m.Ptr())
	return rop1, rop2
}
