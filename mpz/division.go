package mpz

// #include "mpz.h"
import "C"

// Division Functions

//int pmpz_congruent_2exp_p(const unsafe_mpz n, const unsafe_mpz c, mp_bitcnt_t b);
func (m Mpz) MpzCongruent2ExpP(mpz Mpz, bitcnt uint) int {
	ret := C.pmpz_congruent_2exp_p(m.Ptr(), mpz.Ptr(), C.ulong(bitcnt))
	return int(ret)
}

//int pmpz_congruent_p(const unsafe_mpz n, const unsafe_mpz c, const unsafe_mpz d);\
func (m Mpz) MpzCongruentP(c Mpz, d Mpz) int {
	ret := C.pmpz_congruent_p(m.Ptr(), c.Ptr(), d.Ptr())
	return int(ret)
}

//int pmpz_congruent_ui_p(const unsafe_mpz n, unsigned long int c, unsigned long int d);
func (m Mpz) MpzCongruentUiP(c uint, d uint) int {
	ret := C.pmpz_congruent_ui_p(m.Ptr(), C.ulong(c), C.ulong(d))
	return int(ret)
}

//int pmpz_divisible_2exp_p(const unsafe_mpz n, mp_bitcnt_t b);
func (m Mpz) MpzDivisible2ExpP(bitcnt uint) int {
	ret := C.pmpz_divisible_2exp_p(m.Ptr(), C.ulong(bitcnt))
	return int(ret)
}

//int pmpz_divisible_p(const unsafe_mpz n, const unsafe_mpz d);
func (m Mpz) MpzDivisibleP(d Mpz) int {
	ret := C.pmpz_divisible_p(m.Ptr(), d.Ptr())
	return int(ret)
}

//int pmpz_divisible_ui_p(const unsafe_mpz n, unsigned long int d);
func (m Mpz) MpzDivisibleUiP(d uint) int {
	ret := C.pmpz_divisible_ui_p(m.Ptr(), C.ulong(d))
	return int(ret)
}

//unsigned long int pmpz_cdiv_q_ui(unsafe_mpz q, const unsafe_mpz n, unsigned long int d);
func (m Mpz) MpzCDivQUi(d uint) (q Mpz, ui uint) {
	q = MpzInit()
	uiLong := C.pmpz_cdiv_q_ui(q.Ptr(), m.Ptr(), C.ulong(d))
	return q, uint(uiLong)
}

//unsigned long int pmpz_cdiv_qr_ui(unsafe_mpz q, unsafe_mpz r, const unsafe_mpz n, unsigned long int d);
func (m Mpz) MpzCDivQRUi(d uint) (q Mpz, r Mpz, ui uint) {
	r, q = MpzInit(), MpzInit()
	uiLong := C.pmpz_cdiv_qr_ui(q.Ptr(), r.Ptr(), m.Ptr(), C.ulong(d))
	return q, r, uint(uiLong)
}

//unsigned long int pmpz_cdiv_r_ui(unsafe_mpz r, const unsafe_mpz n, unsigned long int d);
func (m Mpz) MpzCDivRUi(d uint) (r Mpz, ui uint) {
	r = MpzInit()
	uiLong := C.pmpz_cdiv_r_ui(r.Ptr(), m.Ptr(), C.ulong(d))
	return r, uint(uiLong)
}

//unsigned long int pmpz_cdiv_ui(const unsafe_mpz n, unsigned long int d);
func (m Mpz) mpzCDivRUi(d uint) (r Mpz, ui uint) {
	r = MpzInit()
	uiLong := C.pmpz_cdiv_r_ui(r.Ptr(), m.Ptr(), C.ulong(d))
	return r, uint(uiLong)
}

//unsigned long int pmpz_fdiv_q_ui(const unsafe_mpz q, const unsafe_mpz n, long unsigned int d);
func (m Mpz) MpzFDivQUi(d uint) (q Mpz, ui uint) {
	q = MpzInit()
	uiLong := C.pmpz_fdiv_q_ui(q.Ptr(), m.Ptr(), C.ulong(d))
	return q, uint(uiLong)
}

//unsigned long int pmpz_fdiv_qr_ui(unsafe_mpz q, unsafe_mpz r, const unsafe_mpz n, long unsigned int d);
func (m Mpz) MpzFDivQRUi(d uint) (q Mpz, r Mpz, ui uint) {
	q, r = MpzInit(), MpzInit()
	uiLong := C.pmpz_fdiv_qr_ui(q.Ptr(), r.Ptr(), m.Ptr(), C.ulong(d))
	return q, r, uint(uiLong)
}

//unsigned long int pmpz_fdiv_r_ui(const unsafe_mpz r, const unsafe_mpz n, long unsigned int d);
func (m Mpz) MpzFDivRUi(d uint) (r Mpz, ui uint) {
	r = MpzInit()
	uiLong := C.pmpz_fdiv_r_ui(r.Ptr(), m.Ptr(), C.ulong(d))
	return r, uint(uiLong)
}

//unsigned long int pmpz_fdiv_ui(const unsafe_mpz n, unsigned long int d);
func (m Mpz) MpzFDivUi(d uint) uint {
	ui := C.pmpz_fdiv_ui(m.Ptr(), C.ulong(d))
	return uint(ui)
}

//unsigned long int pmpz_mod_ui(unsafe_mpz r, const unsafe_mpz n, unsigned long int d);
func (m Mpz) MpzModUi(d uint) (r Mpz, ui uint) {
	r = MpzInit()
	uiLong := C.pmpz_mod_ui(r.Ptr(), m.Ptr(), C.ulong(d))
	return r, uint(uiLong)
}

//unsigned long int pmpz_tdiv_q_ui(unsafe_mpz q, const unsafe_mpz n, unsigned long int d);
func (m Mpz) MpzTDivQUi(d uint) (q Mpz, ui uint) {
	q = MpzInit()
	uiLong := C.pmpz_tdiv_q_ui(q.Ptr(), m.Ptr(), C.ulong(d))
	return q, uint(uiLong)
}

//unsigned long int pmpz_tdiv_qr_ui(unsafe_mpz q, unsafe_mpz r, const unsafe_mpz n, unsigned long int d);.
func (m Mpz) MpzTDivQRUi(d uint) (q Mpz, r Mpz, ui uint) {
	q, r = MpzInit(), MpzInit()
	uiLong := C.pmpz_tdiv_qr_ui(q.Ptr(), r.Ptr(), m.Ptr(), C.ulong(d))
	return q, r, uint(uiLong)
}

//unsigned long int pmpz_tdiv_r_ui(unsafe_mpz r, const unsafe_mpz n, unsigned long int d);
func (m Mpz) MpzTDivRUi(d uint) (r Mpz, ui uint) {
	r = MpzInit()
	uiLong := C.pmpz_tdiv_r_ui(r.Ptr(), m.Ptr(), C.ulong(d))
	return r, uint(uiLong)
}

//unsigned long int pmpz_tdiv_ui(const unsafe_mpz n, unsigned long int d);
func (m Mpz) MpzTDivUi(d uint) uint {
	ret := C.pmpz_tdiv_ui(m.Ptr(), C.ulong(d))
	return uint(ret)
}

//void pmpz_cdiv_qr(unsafe_mpz q, unsafe_mpz r, const unsafe_mpz n, const unsafe_mpz d);
func (m Mpz) MpzCDivQR(d Mpz) (q Mpz, r Mpz) {
	C.pmpz_cdiv_qr(q.Ptr(), r.Ptr(), m.Ptr(), d.Ptr())
	return q, r
}

//void pmpz_fdiv_qr(unsafe_mpz q, unsafe_mpz r, const unsafe_mpz n, const unsafe_mpz d);
func (m Mpz) MpzFDivQR(d Mpz) (q Mpz, r Mpz) {
	C.pmpz_fdiv_qr(q.Ptr(), r.Ptr(), m.Ptr(), d.Ptr())
	return q, r
}

//void pmpz_tdiv_qr(unsafe_mpz q, unsafe_mpz r, const unsafe_mpz n, const unsafe_mpz d);
func (m Mpz) MpzTDivQR(d Mpz) (q Mpz, r Mpz) {
	C.pmpz_tdiv_qr(q.Ptr(), r.Ptr(), m.Ptr(), d.Ptr())
	return q, r
}

//unsafe_mpz pmpz_cdiv_q(const unsafe_mpz n, const unsafe_mpz d);
func (m Mpz) MpzCDivQ(d Mpz) (Mpz, error) {
	ptr := C.pmpz_cdiv_q(m.Ptr(), d.Ptr())
	return mpzFromPtr(ptr)
}

//unsafe_mpz pmpz_cdiv_q_2exp(const unsafe_mpz n, const mp_bitcnt_t b);
func (m Mpz) MpzCDivQ2Exp(bitcnt uint) (Mpz, error) {
	ptr := C.pmpz_cdiv_q_2exp(m.Ptr(), C.ulong(bitcnt))
	return mpzFromPtr(ptr)
}

//unsafe_mpz pmpz_cdiv_r(const unsafe_mpz n, const unsafe_mpz d);
func (m Mpz) MpzCDivR(d Mpz) (Mpz, error) {
	ptr := C.pmpz_cdiv_r(m.Ptr(), d.Ptr())
	return mpzFromPtr(ptr)
}

//unsafe_mpz pmpz_cdiv_r_2exp(const unsafe_mpz n, const mp_bitcnt_t b);
func (m Mpz) MpzCDivR2Exp(bitcnt uint) (Mpz, error) {
	ptr := C.pmpz_cdiv_r_2exp(m.Ptr(), C.ulong(bitcnt))
	return mpzFromPtr(ptr)
}

//unsafe_mpz pmpz_divexact(const unsafe_mpz n, const unsafe_mpz d);
func (m Mpz) MpzCDivExact(d Mpz) (Mpz, error) {
	ptr := C.pmpz_divexact(m.Ptr(), d.Ptr())
	return mpzFromPtr(ptr)
}

//unsafe_mpz pmpz_divexact_ui(const unsafe_mpz n, unsigned long d);
func (m Mpz) MpzCDivExactUi(d C.ulong) (Mpz, error) {
	ptr := C.pmpz_divexact_ui(m.Ptr(), d)
	return mpzFromPtr(ptr)
}

//unsafe_mpz pmpz_fdiv_q(const unsafe_mpz n, const unsafe_mpz d);
func (m Mpz) MpzFDivQ(d Mpz) (Mpz, error) {
	ptr := C.pmpz_fdiv_q(m.Ptr(), d.Ptr())
	return mpzFromPtr(ptr)
}

//unsafe_mpz pmpz_fdiv_q_2exp(const unsafe_mpz n, const mp_bitcnt_t b);
func (m Mpz) MpzFDivQ2Exp(bitcnt C.ulong) (Mpz, error) {
	ptr := C.pmpz_fdiv_q_2exp(m.Ptr(), bitcnt)
	return mpzFromPtr(ptr)
}

//unsafe_mpz pmpz_fdiv_r(const unsafe_mpz n, const unsafe_mpz d);
func (m Mpz) MpzFDivR(d Mpz) (Mpz, error) {
	ptr := C.pmpz_fdiv_r(m.Ptr(), d.Ptr())
	return mpzFromPtr(ptr)
}

//unsafe_mpz pmpz_fdiv_r_2exp(const unsafe_mpz n, const mp_bitcnt_t b);
func (m Mpz) MpzFDivR2Exp(bitcnt C.ulong) (Mpz, error) {
	ptr := C.pmpz_fdiv_r_2exp(m.Ptr(), bitcnt)
	return mpzFromPtr(ptr)
}

//unsafe_mpz pmpz_mod(const unsafe_mpz n, const unsafe_mpz d);
func (m Mpz) MpzMod(bitcnt C.ulong) (Mpz, error) {
	ptr := C.pmpz_fdiv_r_2exp(m.Ptr(), bitcnt)
	return mpzFromPtr(ptr)
}

//unsafe_mpz pmpz_tdiv_q(const unsafe_mpz n, const unsafe_mpz d);
func (m Mpz) MpzTDivQ(d Mpz) (Mpz, error) {
	ptr := C.pmpz_tdiv_q(m.Ptr(), d.Ptr())
	return mpzFromPtr(ptr)
}

//unsafe_mpz pmpz_tdiv_q_2exp(const unsafe_mpz n, const mp_bitcnt_t b);
func (m Mpz) MpzTDivQ2Exp(bitcnt C.ulong) (Mpz, error) {
	ptr := C.pmpz_tdiv_q_2exp(m.Ptr(), bitcnt)
	return mpzFromPtr(ptr)
}

//unsafe_mpz pmpz_tdiv_r(const unsafe_mpz n, const unsafe_mpz d);
func (m Mpz) MpzTDivR(d Mpz) (Mpz, error) {
	ptr := C.pmpz_tdiv_r(m.Ptr(), d.Ptr())
	return mpzFromPtr(ptr)
}

//unsafe_mpz pmpz_tdiv_r_2exp(const unsafe_mpz n, const mp_bitcnt_t b);
func (m Mpz) MpzTDivR2Exp(bitcnt C.ulong) (Mpz, error) {
	ptr := C.pmpz_tdiv_r_2exp(m.Ptr(), bitcnt)
	return mpzFromPtr(ptr)
}
