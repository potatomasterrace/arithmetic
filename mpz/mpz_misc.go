package mpz

// #cgo LDFLAGS: -lgmp
// #include "mpz.h"
import "C"

// Miscellaneous Functions
// int pmpz_even_p(const unsafe_mpz op);
func (op Mpz) MpzEvenP() C.int {
	return C.pmpz_even_p(op.Ptr())
}

// int pmpz_fits_sint_p(const unsafe_mpz op);
func (op Mpz) MpzFitsSintP() C.int {
	return C.pmpz_even_p(op.Ptr())
}

// int pmpz_fits_slong_p(const unsafe_mpz op);
func (op Mpz) MpzFitsSlongP() C.int {
	return C.pmpz_fits_slong_p(op.Ptr())
}

// int pmpz_fits_sshort_p(const unsafe_mpz op);
func (op Mpz) MpzFitsSshortP() C.int {
	return C.pmpz_fits_sshort_p(op.Ptr())
}

// int pmpz_fits_uint_p(const unsafe_mpz op);
func (op Mpz) MpzFitsUintP() C.int {
	return C.pmpz_fits_uint_p(op.Ptr())
}

// int pmpz_fits_ulong_p(const unsafe_mpz op);
func (op Mpz) MpzFitsUlongP() C.int {
	return C.pmpz_fits_ulong_p(op.Ptr())
}

// int pmpz_fits_ushort_p(const unsafe_mpz op);
func (op Mpz) MpzFitsUshortP() C.int {
	return C.pmpz_fits_ushort_p(op.Ptr())
}
