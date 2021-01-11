package mpz

// #include "mpz.h"
import "C"

// OK
// Logical and Bit Manipulation Functions
// int pmpz_tstbit(const unsafe_mpz op, mp_bitcnt_t bit_index);
func (op Mpz) MpzTstBit(bitIndex C.ulong) C.int {
	return C.pmpz_tstbit(op.Ptr(), bitIndex)
}

// mp_bitcnt_t pmpz_hamdist(const unsafe_mpz op1, const unsafe_mpz op2);
func (op1 Mpz) MpzHamDist(op2 Mpz) C.ulong {
	return C.pmpz_hamdist(op1.Ptr(), op2.Ptr())
}

// mp_bitcnt_t pmpz_popcount(const unsafe_mpz op);
func (op Mpz) pmpz_popcount() C.ulong {
	return C.pmpz_popcount(op.Ptr())
}

// mp_bitcnt_t pmpz_scan0(const unsafe_mpz op, mp_bitcnt_t starting_bit);
func (op Mpz) MpzScan0(startingBit C.ulong) C.ulong {
	return C.pmpz_scan0(op.Ptr(), startingBit)
}

// mp_bitcnt_t pmpz_scan1(const unsafe_mpz op, mp_bitcnt_t starting_bit);
func (op Mpz) MpzScan1(startingBit C.ulong) C.ulong {
	return C.pmpz_scan1(op.Ptr(), startingBit)
}

// unsafe_mpz pmpz_and(const unsafe_mpz op1, const unsafe_mpz op2);
func (op1 Mpz) MpzAnd(op2 Mpz) (Mpz, error) {
	ptr := C.pmpz_and(op1.Ptr(), op2.Ptr())
	return mpzFromPtr(ptr)
}

// unsafe_mpz pmpz_com(const unsafe_mpz op);
func (op Mpz) MpzCom() (Mpz, error) {
	ptr := C.pmpz_com(op.Ptr())
	return mpzFromPtr(ptr)
}

// unsafe_mpz pmpz_ior(const unsafe_mpz op1, const unsafe_mpz op2);
func (op1 Mpz) MpzIOr(op2 Mpz) (Mpz, error) {
	ptr := C.pmpz_ior(op1.Ptr(), op2.Ptr())
	return mpzFromPtr(ptr)
}

// unsafe_mpz pmpz_xor(const unsafe_mpz op1, const unsafe_mpz op2);
func (op1 Mpz) MpzXor(op2 Mpz) (Mpz, error) {
	ptr := C.pmpz_xor(op1.Ptr(), op2.Ptr())
	return mpzFromPtr(ptr)
}

// TODO add variadic variations of these functions

// void pmpz_clrbit(unsafe_mpz rop, mp_bitcnt_t bit_index);
func (op Mpz) MpzClrBit(bitIndex C.ulong) (Mpz, error) {
	rop, err := MpzFromMpz(op)
	if err != nil {
		return rop, err
	}
	C.pmpz_clrbit(rop.Ptr(), bitIndex)
	return rop, nil
}

// void pmpz_combit(unsafe_mpz rop, mp_bitcnt_t bit_index);
func (op Mpz) MpzComBit(bitIndex C.ulong) (Mpz, error) {
	rop, err := MpzFromMpz(op)
	if err != nil {
		return rop, err
	}
	C.pmpz_combit(rop.Ptr(), bitIndex)
	return rop, nil
}

// void pmpz_setbit(const unsafe_mpz op, mp_bitcnt_t bit_index);
func (op Mpz) MpzSetBit(bitIndex C.ulong) (Mpz, error) {
	rop, err := MpzFromMpz(op)
	if err != nil {
		return rop, err
	}
	C.pmpz_combit(rop.Ptr(), bitIndex)
	return rop, nil
}
