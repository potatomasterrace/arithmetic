package mpz

// #include "mpz.h"
import "C"

// Comparison Functions
// int pmpz_cmp(const unsafe_mpz op1, const unsafe_mpz op2);
func (m Mpz) Cmp(mpz Mpz) int {
	ret := C.pmpz_cmp(m.Ptr(), mpz.Ptr())
	return int(ret)
}

// int pmpz_cmp_d(const unsafe_mpz op1, double op2);
func (m Mpz) CmpD(val int) int {
	ret := C.pmpz_cmp_d(m.Ptr(), C.double(val))
	return int(ret)
}

// int pmpz_cmp_si(const unsafe_mpz op1, signed long int op2);
func (m Mpz) CmpSi(val int) int {
	ret := C.pmpz_cmp_si(m.Ptr(), C.long(val))
	return int(ret)
}

// int pmpz_cmp_ui(const unsafe_mpz op1, unsigned long int op2);
func (m Mpz) CmpUi(val uint) int {
	ret := C.pmpz_cmp_ui(m.Ptr(), C.ulong(val))
	return int(ret)
}

// int pmpz_cmpabs(const unsafe_mpz op1, const unsafe_mpz op2);
func (m Mpz) CmpAbs(mpz Mpz) int {
	ret := C.pmpz_cmpabs(m.Ptr(), mpz.Ptr())
	return int(ret)
}

// int pmpz_cmpabs_d(const unsafe_mpz op1, double op2);
func (m Mpz) CmpAbsD(val float64) int {
	ret := C.pmpz_cmpabs_d(m.Ptr(), C.double(val))
	return int(ret)
}

// int pmpz_cmpabs_ui(const unsafe_mpz op1, unsigned long int op2);
func (m Mpz) CmpAbsUi(val uint) int {
	ret := C.pmpz_cmpabs_ui(m.Ptr(), C.ulong(val))
	return int(ret)
}

// int pmpz_sgn(const unsafe_mpz op);
func (m Mpz) Sign() int {
	ret := C.pmpz_sgn(m.Ptr())
	return int(ret)
}
