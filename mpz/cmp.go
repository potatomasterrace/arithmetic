package mpz

// #include "mpz.h"
import "C"

// Comparison Functions

// Cmp compares op1 and op2. Return a positive value if op1 > op2, zero if op1 = op2, or a negative value if op1 < op2.
// underlying gmp function : mpz_cmp
func (op1 Mpz) Cmp(op2 Mpz) int {
	ret := C.pmpz_cmp(op1.Ptr(), op2.Ptr())
	return int(ret)
}

// CmpD compares op1 and op2. Return a positive value if op1 > op2, zero if op1 = op2, or a negative value if op1 < op2.
// underlying gmp function : mpz_cmp_d
func (op1 Mpz) CmpD(op2 float64) int {
	ret := C.pmpz_cmp_d(op1.Ptr(), C.double(op2))
	return int(ret)
}

// CmpSi compares op1 and op2. Return a positive value if op1 > op2, zero if op1 = op2, or a negative value if op1 < op2.
// underlying gmp function : mpz_cmp_si
func (op1 Mpz) CmpSi(op2 int) int {
	ret := C.pmpz_cmp_si(op1.Ptr(), C.long(op2))
	return int(ret)
}

// CmpUi compares op1 and op2. Return a positive value if op1 > op2, zero if op1 = op2, or a negative value if op1 < op2.
// underlying gmp function : mpz_cmp_ui
func (op1 Mpz) CmpUi(op2 uint) int {
	ret := C.pmpz_cmp_ui(op1.Ptr(), C.ulong(op2))
	return int(ret)
}

// CmpAbs compares op1 and op2. Return a positive value if op1 > op2, zero if op1 = op2, or a negative value if op1 < op2.
// underlying gmp function : mpz_cmpabs
func (op1 Mpz) CmpAbs(op2 Mpz) int {
	ret := C.pmpz_cmpabs(op1.Ptr(), op2.Ptr())
	return int(ret)
}

// CmpAbsD compares op1 and op2. Return a positive value if op1 > op2, zero if op1 = op2, or a negative value if op1 < op2.
// underlying gmp function : mpz_cmpabs_d
func (op1 Mpz) CmpAbsD(op2 float64) int {
	ret := C.pmpz_cmpabs_d(op1.Ptr(), C.double(op2))
	return int(ret)
}

// CmpAbsUi compares op1 and op2. Return a positive value if op1 > op2, zero if op1 = op2, or a negative value if op1 < op2.
// underlying gmp function : mpz_cmpabs_ui
func (op1 Mpz) CmpAbsUi(op2 uint) int {
	ret := C.pmpz_cmpabs_ui(op1.Ptr(), C.ulong(op2))
	return int(ret)
}

// Sign compares op1 and op2. Return a positive value if op1 > op2, zero if op1 = op2, or a negative value if op1 < op2.
// underlying gmp function : mpz_cmp_d
func (op Mpz) Sign() int {
	ret := C.pmpz_sgn(op.Ptr())
	return int(ret)
}
