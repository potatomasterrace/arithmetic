package mpz

// #include "mpz.h"
import "C"

// Comparison Functions

// Cmp compares op1 and op2. Return a positive value if op1 > op2, zero if op1 = op2, or a negative value if op1 < op2.
// underlying gmp function : mpz_cmp
func (op1 Mpz) Cmp(op2 Mpz) (int, error) {
	p1, p2, err := get2Pointers(op1, op2)
	if err != nil {
		return 0, err
	}
	ret := C.pmpz_cmp(p1, p2)
	return int(ret), nil
}

func (op1 Mpz) CmpOrPanic(op2 Mpz) int {
	ret, err := op1.Cmp(op2)
	if err != nil {
		panic(err)
	}
	return ret
}

// CmpD compares op1 and op2. Return a positive value if op1 > op2, zero if op1 = op2, or a negative value if op1 < op2.
// underlying gmp function : mpz_cmp_d
func (op1 Mpz) CmpD(op2 float64) (int, error) {
	p1, err := op1.Ptr()
	if err != nil {
		return 0, err
	}
	ret := C.pmpz_cmp_d(p1, C.double(op2))
	return int(ret), nil
}

func (op1 Mpz) CmpDOrPanic(op2 float64) int {
	ret, err := op1.CmpD(op2)
	if err != nil {
		panic(err)
	}
	return ret
}

// CmpSi compares op1 and op2. Return a positive value if op1 > op2, zero if op1 = op2, or a negative value if op1 < op2.
// underlying gmp function : mpz_cmp_si
func (op1 Mpz) CmpSi(op2 int) (int, error) {
	p1, err := op1.Ptr()
	if err != nil {
		return 0, err
	}
	ret := C.pmpz_cmp_si(p1, C.long(op2))
	return int(ret), nil
}

func (op1 Mpz) CmpSiOrPanic(op2 int) int {
	ret, err := op1.CmpSi(op2)
	if err != nil {
		panic(err)
	}
	return ret
}

// CmpUi compares op1 and op2. Return a positive value if op1 > op2, zero if op1 = op2, or a negative value if op1 < op2.
// underlying gmp function : mpz_cmp_ui
func (op1 Mpz) CmpUi(op2 uint) (int, error) {
	p1, err := op1.Ptr()
	if err != nil {
		return 0, err
	}
	ret := C.pmpz_cmp_ui(p1, C.ulong(op2))
	return int(ret), nil
}

func (op1 Mpz) CmpUiOrPanic(op2 uint) int {
	ret, err := op1.CmpUi(op2)
	if err != nil {
		panic(err)
	}
	return ret
}

// CmpAbs compares op1 and op2. Return a positive value if op1 > op2, zero if op1 = op2, or a negative value if op1 < op2.
// underlying gmp function : mpz_cmpabs
func (op1 Mpz) CmpAbs(op2 Mpz) (int, error) {
	p1, p2, err := get2Pointers(op1, op2)
	if err != nil {
		return 0, err
	}
	ret := C.pmpz_cmpabs(p1, p2)
	return int(ret), nil
}

func (op1 Mpz) CmpAbsOrPanic(op2 Mpz) int {
	ret, err := op1.CmpAbs(op2)
	if err != nil {
		panic(err)
	}
	return ret
}

// CmpAbsD compares op1 and op2. Return a positive value if op1 > op2, zero if op1 = op2, or a negative value if op1 < op2.
// underlying gmp function : mpz_cmpabs_d
func (op1 Mpz) CmpAbsD(op2 float64) (int, error) {
	p1, err := op1.Ptr()
	if err != nil {
		return 0, err
	}
	ret := C.pmpz_cmpabs_d(p1, C.double(op2))
	return int(ret), nil
}

func (op1 Mpz) CmpAbsDOrPanic(op2 float64) int {
	ret, err := op1.CmpAbsD(op2)
	if err != nil {
		panic(err)
	}
	return ret
}

// CmpAbsUi compares op1 and op2. Return a positive value if op1 > op2, zero if op1 = op2, or a negative value if op1 < op2.
// underlying gmp function : mpz_cmpabs_ui
func (op1 Mpz) CmpAbsUi(op2 uint) (int, error) {
	p1, err := op1.Ptr()
	if err != nil {
		return 0, err
	}
	ret := C.pmpz_cmpabs_ui(p1, C.ulong(op2))
	return int(ret), nil
}

func (op1 Mpz) CmpAbsUiOrPanic(op2 uint) int {
	ret, err := op1.CmpAbsUi(op2)
	if err != nil {
		panic(err)
	}
	return ret
}

// Sign compares op1 and op2. Return a positive value if op1 > op2, zero if op1 = op2, or a negative value if op1 < op2.
// underlying gmp function : mpz_cmp_d
func (op Mpz) Sign() (int, error) {
	p, err := op.Ptr()
	if err != nil {
		return 0, err
	}
	ret := C.pmpz_sgn(p)
	return int(ret), nil
}

func (op Mpz) SignOrPanic() int {
	p, err := op.Ptr()
	if err != nil {
		panic(err)
	}
	ret := C.pmpz_sgn(p)
	return int(ret)
}
