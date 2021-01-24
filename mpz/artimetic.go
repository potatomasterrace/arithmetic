package mpz

// #include "mpz.h"
import "C"
import "unsafe"

// Arithmetic Functions
// Given in the same order than the documentation here : https://gmplib.org/manual/Integer-Arithmetic

// Add op1 to op2 and return the result.
// underlying gmp function : mpz_add
func (op1 Mpz) Add(op2 Mpz) (Mpz, error) {
	p1, p2, err := get2Pointers(op1, op2)
	if err != nil {
		return Mpz{}, err
	}
	ptr := C.pmpz_add(p1, p2)
	return mpzFromPtr(ptr)
}

// AddMany adds all the values passed to op1.
// underlying gmp function : mpz_add
func (op1 Mpz) AddMany(ops ...Mpz) (Mpz, error) {
	sum := op1
	for _, op := range ops {
		var err error
		sum, err = sum.Add(op)
		if err != nil {
			return sum, err
		}
	}
	return sum, nil
}

// Add op1 to op2 and return the result.
// underlying gmp function : mpz_add_ui
func (op1 Mpz) AddUi(op2 uint) (Mpz, error) {
	p1, err := op1.Ptr()
	if err != nil {
		return Mpz{}, err
	}
	ptr := C.pmpz_add_ui(p1, C.ulong(op2))
	return mpzFromPtr(ptr)
}

// Substract op1 to op2 and return the result.
// underlying gmp function : mpz_sub
func (op1 Mpz) Sub(op2 Mpz) (Mpz, error) {
	p1, p2, err := get2Pointers(op1, op2)
	if err != nil {
		return Mpz{}, err
	}
	ptr := C.pmpz_sub(p1, p2)
	return mpzFromPtr(ptr)
}

// SubMany substracts all the values passed to op1.
// underlying gmp function : mpz_sub
func (op1 Mpz) SubMany(ops ...Mpz) (Mpz, error) {
	sub := op1
	for _, op := range ops {
		var err error
		sub, err = sub.Sub(op)
		if err != nil {
			return sub, err
		}
	}
	return sub, nil
}

// Substract op1 to op2 and return the result.
// underlying gmp function : mpz_sub_ui
func (op1 Mpz) SubUi(op2 uint) (Mpz, error) {
	p1, err := op1.Ptr()
	if err != nil {
		return Mpz{}, err
	}
	ptr := C.pmpz_sub_ui(p1, C.ulong(op2))
	return mpzFromPtr(ptr)
}

// Substract op1 to op2 and return the result.
// underlying gmp function : mpz_ui_sub
func (op2 Mpz) UiSub(op1 uint) (Mpz, error) {
	p2, err := op2.Ptr()
	if err != nil {
		return Mpz{}, err
	}
	ptr := C.pmpz_ui_sub(C.ulong(op1), p2)
	return mpzFromPtr(ptr)
}

// Multiply op1 times op2 and return the result.
// underlying gmp function : mpz_mul
func (op1 Mpz) Mul(op2 Mpz) (Mpz, error) {
	p1, p2, err := get2Pointers(op1, op2)
	if err != nil {
		return Mpz{}, err
	}
	ptr := C.pmpz_mul(p1, p2)
	return mpzFromPtr(ptr)
}

// MulMany multiplies all the values passed to op1.
// underlying gmp function : mpz_sub
func (op1 Mpz) MulMany(ops ...Mpz) (Mpz, error) {
	prod := op1
	for _, op := range ops {
		var err error
		prod, err = prod.Mul(op)
		if err != nil {
			return prod, err
		}
	}
	return prod, nil
}

// Multiply op1 times op2 and return the result.
// underlying gmp function : mpz_mul_si
func (op1 Mpz) MulSi(op2 int) (Mpz, error) {
	p1, err := op1.Ptr()
	if err != nil {
		return Mpz{}, err
	}
	ptr := C.pmpz_mul_si(p1, C.long(op2))
	return mpzFromPtr(ptr)
}

// Multiply op1 times op2 and return the result.
// underlying gmp function : mpz_mul_ui
func (op1 Mpz) MulUi(op2 uint) (Mpz, error) {
	p1, err := op1.Ptr()
	if err != nil {
		return Mpz{}, err
	}
	ptr := C.pmpz_mul_ui(p1, C.ulong(op2))
	return mpzFromPtr(ptr)
}

// AddMul returns m+op1 times op2 .
// underlying gmp function : mpz_unsafe_addmul
func (m Mpz) AddMul(op1 Mpz, op2 Mpz) (Mpz, error) {
	mCopy, err := MpzFromMpz(m)
	if err != nil {
		return mCopy, err
	}
	p1, p2, err := get2Pointers(op1, op2)
	if err != nil {
		return Mpz{}, err
	}
	mCopyPtr, _ := mCopy.Ptr()
	C.pmpz_unsafe_addmul(mCopyPtr, p1, p2)
	return mCopy, nil
}

// AddMulUi returns m+op1 times op2 .
// underlying gmp function : mpz_unsafe_addmul_ui
func (m Mpz) AddMulUi(op1 Mpz, op2 uint) (Mpz, error) {
	mCopy, err := MpzFromMpz(m)
	if err != nil {
		return mCopy, err
	}
	p1, err := op1.Ptr()
	if err != nil {
		return Mpz{}, err
	}
	mCopyPtr, _ := mCopy.Ptr()
	C.pmpz_unsafe_addmul_ui(mCopyPtr, p1, C.ulong(op2))
	return mCopy, nil
}

// SubMul returns m-op1 times op2 .
// underlying gmp function : mpz_unsafe_submul
func (m Mpz) SubMul(op1 Mpz, op2 Mpz) (Mpz, error) {
	mCopy, err := MpzFromMpz(m)
	if err != nil {
		return mCopy, err
	}
	p1, p2, err := get2Pointers(op1, op2)
	if err != nil {
		return Mpz{}, err
	}
	mCopyPtr, _ := mCopy.Ptr()
	C.pmpz_unsafe_submul(mCopyPtr, p1, p2)
	return mCopy, nil
}

// SubMulUi returns m-op1 times op2 .
// underlying gmp function : mpz_unsafe_submul_ui
func (m Mpz) SubMulUi(op1 Mpz, op2 uint) (Mpz, error) {
	mCopy, err := MpzFromMpz(m)
	if err != nil {
		return mCopy, err
	}
	p1, err := op1.Ptr()
	if err != nil {
		return Mpz{}, err
	}
	mCopyPtr, _ := mCopy.Ptr()
	C.pmpz_unsafe_submul_ui(mCopyPtr, p1, C.ulong(op2))
	return mCopy, nil
}

// Mul2Exp returns m-op1 times op2 .
// underlying gmp function : mpz_mul_2exp
func (op1 Mpz) Mul2Exp(op2 uint) (Mpz, error) {
	p1, err := op1.Ptr()
	if err != nil {
		return Mpz{}, err
	}
	ptr := C.pmpz_mul_2exp(p1, C.ulong(op2))
	return mpzFromPtr(ptr)
}

// Neg return the negative.
// underlying gmp function : mpz_neg
func (op Mpz) Neg() (Mpz, error) {
	p, err := op.Ptr()
	if err != nil {
		return Mpz{}, err
	}
	ptr := C.pmpz_neg(p)
	return mpzFromPtr(ptr)
}

// Abs return the absolute value.
// underlying gmp function : mpz_abs
func (op Mpz) Abs() (Mpz, error) {
	p, err := op.Ptr()
	if err != nil {
		return Mpz{}, err
	}
	ptr := C.pmpz_abs(p)
	return mpzFromPtr(ptr)
}

func get2Pointers(op1 Mpz, op2 Mpz) (unsafe.Pointer, unsafe.Pointer, error) {
	p1, err := op1.Ptr()
	if err != nil {
		return nil, nil, err
	}
	p2, err := op2.Ptr()
	if err != nil {
		return nil, nil, err
	}
	return p1, p2, nil
}

func get3Pointers(op1 Mpz, op2 Mpz, op3 Mpz) (unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, error) {
	p1, p2, err := get2Pointers(op1, op2)
	if err != nil {
		return nil, nil, nil, err
	}
	p3, err := op3.Ptr()
	if err != nil {
		return nil, nil, nil, err
	}
	return p1, p2, p3, nil
}

func get4Pointers(op1 Mpz, op2 Mpz, op3 Mpz, op4 Mpz) (unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, error) {
	p1, p2, err := get2Pointers(op1, op2)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	p3, p4, err := get2Pointers(op3, op4)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	return p1, p2, p3, p4, nil
}
