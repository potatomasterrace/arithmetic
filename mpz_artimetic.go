package main

// #cgo LDFLAGS: -lgmp
// #include "mpz.h"
import "C"

// Arithmetic Functions
// unsafe_mpz pmpz_abs(const unsafe_mpz op);
func (m Mpz) AbsMpz() Mpz {
	ptr := C.pmpz_abs(m.Ptr())
	return mpzFromPtr(ptr)
}

// unsafe_mpz pmpz_add(const unsafe_mpz op1, const unsafe_mpz op2);
func (m Mpz) AddMpz(mpz Mpz) Mpz {
	ptr := C.pmpz_add(m.Ptr(), mpz.Ptr())
	return mpzFromPtr(ptr)
}

// unsafe_mpz pmpz_add_ui(const unsafe_mpz op1, const unsigned long int op2);
func (m Mpz) AddUiMpz(val C.ulong) Mpz {
	ptr := C.pmpz_add_ui(m.Ptr(), val)
	return mpzFromPtr(ptr)
}

// unsafe_mpz pmpz_mul(const unsafe_mpz op1, const unsafe_mpz op2);
func (m Mpz) MulMpz(op1 Mpz, op2 Mpz) Mpz {
	ptr := C.pmpz_mul(op1.Ptr(), op2.Ptr())
	return mpzFromPtr(ptr)
}

// unsafe_mpz pmpz_mul_2exp(const unsafe_mpz op1, const mp_bitcnt_t op2);
func (m Mpz) Mul2ExpMpz(val C.ulong) Mpz {
	ptr := C.pmpz_mul_2exp(m.Ptr(), val)
	return mpzFromPtr(ptr)
}

// unsafe_mpz pmpz_mul_si(const unsafe_mpz op1, long int op2);
func (m Mpz) MultSiMpz(val C.long) Mpz {
	ptr := C.pmpz_mul_si(m.Ptr(), val)
	return mpzFromPtr(ptr)
}

// unsafe_mpz pmpz_mul_ui(const unsafe_mpz op1, unsigned long int op2);
func (m Mpz) MultUIMpz(val C.ulong) Mpz {
	ptr := C.pmpz_mul_ui(m.Ptr(), val)
	return mpzFromPtr(ptr)
}

// unsafe_mpz pmpz_neg(const unsafe_mpz op);
func (m Mpz) NegMpz() Mpz {
	ptr := C.pmpz_neg(m.Ptr())
	return mpzFromPtr(ptr)
}

// unsafe_mpz pmpz_sub(const unsafe_mpz op1, const unsafe_mpz op2);
func (m Mpz) SubMpz(mpz Mpz) Mpz {
	ptr := C.pmpz_sub(m.Ptr(), mpz.Ptr())
	return mpzFromPtr(ptr)
}

// unsafe_mpz pmpz_sub_ui(const unsafe_mpz op1, const unsigned long int op2);
func (m Mpz) SubUiMpz(val C.ulong) Mpz {
	ptr := C.pmpz_sub_ui(m.Ptr(), val)
	return mpzFromPtr(ptr)
}

// unsafe_mpz pmpz_ui_sub(const unsigned long int op1, const unsafe_mpz op2);
func (m Mpz) UiSubMpz(val C.ulong) Mpz {
	ptr := C.pmpz_ui_sub(val, m.Ptr())
	return mpzFromPtr(ptr)
}

// unsafe_mpz pmpz_unsafe_addmul(const unsafe_mpz rop, const unsafe_mpz op1, const unsigned long int op2);
func (m Mpz) UnsafeAddMulMpz(op1 Mpz, op2 Mpz) Mpz {
	C.pmpz_unsafe_addmul(m.Ptr(), op1.Ptr(), op2.Ptr())
	return m
}

func (m Mpz) AddMulMpz(op1 Mpz, op2 Mpz) Mpz {
	mCopy := MpzFromMpz(m)
	return mCopy.UnsafeAddMulMpz(op1, op2)
}

// unsafe_mpz pmpz_unsafe_addmul_ui(const unsafe_mpz rop, const unsafe_mpz op1, const unsigned long int op2);
func (m Mpz) UnsafeAddMulUiMpz(op1 Mpz, op2 C.ulong) Mpz {
	C.pmpz_unsafe_addmul_ui(m.Ptr(), op1.Ptr(), op2)
	return m
}
func (m Mpz) AddMulUI(op1 Mpz, op2 C.ulong) Mpz {
	mCopy := MpzFromMpz(m)
	return mCopy.UnsafeAddMulUiMpz(op1, op2)
}

// unsafe_mpz pmpz_unsafe_submul(const unsafe_mpz rop, const unsafe_mpz op1, const unsafe_mpz op2);
func (m Mpz) UnsafeSubMulMpz(op1 Mpz, op2 Mpz) Mpz {
	C.pmpz_unsafe_submul(m.Ptr(), op1.Ptr(), op2.Ptr())
	return m
}
func (m Mpz) SubMul(op1 Mpz, op2 Mpz) Mpz {
	mCopy := MpzFromMpz(m)
	return mCopy.UnsafeSubMulMpz(op1, op2)
}

// unsafe_mpz pmpz_unsafe_submul_ui(const unsafe_mpz rop, const unsafe_mpz op1, const unsigned long int op2);
func (m Mpz) UnsafeSubMulUiMpz(op1 Mpz, op2 C.ulong) Mpz {
	C.pmpz_unsafe_submul_ui(m.Ptr(), op1.Ptr(), op2)
	return m
}

func (m Mpz) SubMulUiMpz(op1 Mpz, op2 C.ulong) Mpz {
	mCopy := MpzFromMpz(m)
	return mCopy.UnsafeSubMulUiMpz(op1, op2)
}
