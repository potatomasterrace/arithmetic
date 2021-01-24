package mpz

// #include "mpz.h"
import "C"

// Arithmetic Functions
// Given in the same order than the documentation here : https://gmplib.org/manual/Integer-Arithmetic

// Add op1 to op2 and return the result.
// underlying gmp function : mpz_add
func (op1 Mpz) Add(op2 Mpz) Mpz {
	ptr := C.pmpz_add(op1.Ptr(), op2.Ptr())
	return mpzFromPtr(ptr)
}

// AddMany adds all the values passed to op1.
// underlying gmp function : mpz_add
func (op1 Mpz) AddMany(ops ...Mpz) Mpz {
	sum := op1
	for _, op := range ops {
		sum = sum.Add(op)
	}
	return sum
}

// Add op1 to op2 and return the result.
// underlying gmp function : mpz_add_ui
func (op1 Mpz) AddUi(op2 uint) Mpz {
	ptr := C.pmpz_add_ui(op1.Ptr(), C.ulong(op2))
	return mpzFromPtr(ptr)
}

// Substract op1 to op2 and return the result.
// underlying gmp function : mpz_sub
func (op1 Mpz) Sub(op2 Mpz) Mpz {
	ptr := C.pmpz_sub(op1.Ptr(), op2.Ptr())
	return mpzFromPtr(ptr)
}

// SubMany substracts all the values passed to op1.
// underlying gmp function : mpz_sub
func (op1 Mpz) SubMany(ops ...Mpz) Mpz {
	sub := op1
	for _, op := range ops {
		sub = sub.Sub(op)
	}
	return sub
}

// Substract op1 to op2 and return the result.
// underlying gmp function : mpz_sub_ui
func (op1 Mpz) SubUi(op2 uint) Mpz {
	ptr := C.pmpz_sub_ui(op1.Ptr(), C.ulong(op2))
	return mpzFromPtr(ptr)
}

// Substract op1 to op2 and return the result.
// underlying gmp function : mpz_ui_sub
func (op2 Mpz) UiSub(op1 uint) Mpz {
	ptr := C.pmpz_ui_sub(C.ulong(op1), op2.Ptr())
	return mpzFromPtr(ptr)
}

// Multiply op1 times op2 and return the result.
// underlying gmp function : mpz_mul
func (op1 Mpz) Mul(op2 Mpz) Mpz {
	ptr := C.pmpz_mul(op1.Ptr(), op2.Ptr())
	return mpzFromPtr(ptr)
}

// MulMany multiplies all the values passed to op1.
// underlying gmp function : mpz_sub
func (op1 Mpz) MulMany(ops ...Mpz) Mpz {
	prod := op1
	for _, op := range ops {
		prod = prod.Mul(op)
	}
	return prod
}

// Multiply op1 times op2 and return the result.
// underlying gmp function : mpz_mul_si
func (op1 Mpz) MulSi(op2 int) Mpz {
	ptr := C.pmpz_mul_si(op1.Ptr(), C.long(op2))
	return mpzFromPtr(ptr)
}

// Multiply op1 times op2 and return the result.
// underlying gmp function : mpz_mul_ui
func (op1 Mpz) MulUi(op2 uint) Mpz {
	ptr := C.pmpz_mul_ui(op1.Ptr(), C.ulong(op2))
	return mpzFromPtr(ptr)
}

// AddMul returns m+op1 times op2 .
// underlying gmp function : mpz_unsafe_addmul
func (m Mpz) AddMul(op1 Mpz, op2 Mpz) Mpz {
	mCopy := MpzFromMpz(m)
	C.pmpz_unsafe_addmul(mCopy.Ptr(), op1.Ptr(), op2.Ptr())
	return mCopy
}

// AddMulUi returns m+op1 times op2 .
// underlying gmp function : mpz_unsafe_addmul_ui
func (m Mpz) AddMulUi(op1 Mpz, op2 uint) Mpz {
	mCopy := MpzFromMpz(m)
	C.pmpz_unsafe_addmul_ui(mCopy.Ptr(), op1.Ptr(), C.ulong(op2))
	return mCopy
}

// SubMul returns m-op1 times op2 .
// underlying gmp function : mpz_unsafe_submul
func (m Mpz) SubMul(op1 Mpz, op2 Mpz) Mpz {
	mCopy := MpzFromMpz(m)
	C.pmpz_unsafe_submul(mCopy.Ptr(), op1.Ptr(), op2.Ptr())
	return mCopy
}

// SubMulUi returns m-op1 times op2 .
// underlying gmp function : mpz_unsafe_submul_ui
func (m Mpz) SubMulUi(op1 Mpz, op2 uint) Mpz {
	mCopy := MpzFromMpz(m)

	C.pmpz_unsafe_submul_ui(mCopy.Ptr(), op1.Ptr(), C.ulong(op2))
	return mCopy
}

// Mul2Exp returns m-op1 times op2 .
// underlying gmp function : mpz_mul_2exp
func (op1 Mpz) Mul2Exp(op2 uint) Mpz {
	ptr := C.pmpz_mul_2exp(op1.Ptr(), C.ulong(op2))
	return mpzFromPtr(ptr)
}

// Neg return the negative.
// underlying gmp function : mpz_neg
func (op Mpz) Neg() Mpz {
	ptr := C.pmpz_neg(op.Ptr())
	return mpzFromPtr(ptr)
}

// Abs return the absolute value.
// underlying gmp function : mpz_abs
func (op Mpz) Abs() Mpz {
	ptr := C.pmpz_abs(op.Ptr())
	return mpzFromPtr(ptr)
}
