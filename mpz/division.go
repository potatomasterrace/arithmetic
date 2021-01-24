package mpz

// #include "mpz.h"
import "C"

// Division Functions

// Congruent2ExpP returns if n is congruent to c modulo 2^b.
// underlying gmp function : mpz_congruent_2exp_p
func (n Mpz) Congruent2ExpP(c Mpz, b uint) (bool, error) {
	nPtr, cPtr, err := get2Pointers(n, c)
	if err != nil {
		return false, err
	}
	ret := C.pmpz_congruent_2exp_p(nPtr, cPtr, C.ulong(b))
	return ret != 0, nil
}

func (n Mpz) Congruent2ExpPOrPanic(c Mpz, b uint) bool {
	nPtr, cPtr, err := get2Pointers(n, c)
	if err != nil {
		panic(err)
	}
	ret := C.pmpz_congruent_2exp_p(nPtr, cPtr, C.ulong(b))
	return ret != 0
}

// CongruentP returns if n is congruent to c modulo d.
// underlying gmp function : mpz_congruent_p
func (n Mpz) CongruentP(c Mpz, d Mpz) (bool, error) {
	nPtr, cPtr, dPtr, err := get3Pointers(n, c, d)
	if err != nil {
		return false, err
	}
	ret := C.pmpz_congruent_p(nPtr, cPtr, dPtr)
	return ret != 0, nil
}

func (n Mpz) CongruentPOrPanic(c Mpz, d Mpz) bool {
	nPtr, cPtr, dPtr, err := get3Pointers(n, c, d)
	if err != nil {
		panic(err)
	}
	ret := C.pmpz_congruent_p(nPtr, cPtr, dPtr)
	return ret != 0
}

// CongruentUiP returns if n is congruent to c modulo d.
// underlying gmp function : mpz_congruent_ui_p
func (n Mpz) CongruentUiP(c uint, d uint) (bool, error) {
	nPtr, err := n.Ptr()
	if err != nil {
		return false, err
	}
	ret := C.pmpz_congruent_ui_p(nPtr, C.ulong(c), C.ulong(d))
	return ret != 0, nil
}

// Divisible2ExpP returns if n is divisible by 2^b.
// underlying gmp function : mpz_divisible_2exp_p
func (n Mpz) Divisible2ExpP(b uint) (bool, error) {
	nPtr, err := n.Ptr()
	if err != nil {
		return false, err
	}
	ret := C.pmpz_divisible_2exp_p(nPtr, C.ulong(b))
	return ret != 0, nil
}

// DivisibleP return if n is divisible by d.
// underlying gmp function : mpz_divisible_p
func (n Mpz) DivisibleP(d Mpz) (bool, error) {
	nPtr, dPtr, err := get2Pointers(n, d)
	if err != nil {
		return false, err
	}
	ret := C.pmpz_divisible_p(nPtr, dPtr)
	return ret != 0, nil
}

// DivisibleUiP returns if n is divisible by d.
// underlying gmp function : mpz_divisible_p
func (n Mpz) DivisibleUiP(d uint) (bool, error) {
	nPtr, err := n.Ptr()
	if err != nil {
		return false, err
	}
	ret := C.pmpz_divisible_ui_p(nPtr, C.ulong(d))
	return ret != 0, nil
}

// CDivQUi returns q = floor(n / d) and  r = n - d * q
// underlying gmp function : mpz_cdiv_q_ui
func (n Mpz) CDivQUi(d uint) (q Mpz, r uint, err error) {
	q = MpzInit()
	qPtr, nPtr, err := get2Pointers(q, n)
	if err != nil {
		return q, 0, err
	}
	uiLong := C.pmpz_cdiv_q_ui(qPtr, nPtr, C.ulong(d))
	return q, uint(uiLong), nil
}

// CDivQRUi returns q = floor(n / d) and  r = n - d * q
// the remainder is also return as an ui if it fits an uint.
// underlying gmp function : mpz_cdiv_q_ui
func (n Mpz) CDivQRUi(d uint) (q Mpz, r Mpz, rUi uint, err error) {
	r, q = MpzInit(), MpzInit()
	qPtr, rPtr, nPtr, err := get3Pointers(q, r, n)
	if err != nil {
		return q, r, 0, err
	}
	ret := C.pmpz_cdiv_qr_ui(qPtr, rPtr, nPtr, C.ulong(d))
	return q, r, uint(ret), nil
}

// CDivRUi returns the remainder r = | n - q * d |,
// the remainder is also return as an ui if it fits an uint.
// underlying gmp function : mpz_cdiv_r_ui
func (n Mpz) CDivRUi(d uint) (r Mpz, rUi uint, err error) {
	r = MpzInit()
	rPtr, nPtr, err := get2Pointers(r, n)
	if err != nil {
		return r, 0, err
	}
	ret := C.pmpz_cdiv_r_ui(rPtr, nPtr, C.ulong(d))
	return r, uint(ret), err
}

// CDivRUi returns the remainder r = | n - q * d |,
// the remainder is also return as an ui if it fits an uint.
// underlying gmp function : mpz_cdiv_ui
func (n Mpz) CDivUi(d uint) (r uint, err error) {
	nPtr, err := n.Ptr()
	if err != nil {
		return 0, err
	}
	ret := C.pmpz_cdiv_ui(nPtr, C.ulong(d))
	return uint(ret), err
}

// FDivQUi returns q = floor(n / d) and  r = n - d * q
// underlying gmp function : pmpz_fdiv_q_ui
func (n Mpz) FDivQUi(d uint) (q Mpz, r uint, err error) {
	ret := MpzInit()
	retPtr, nPtr, err := get2Pointers(ret, n)
	if err != nil {
		return ret, 0, err
	}
	retUi := C.pmpz_fdiv_q_ui(retPtr, nPtr, C.ulong(d))
	return ret, uint(retUi), err
}

//unsigned long int pmpz_fdiv_qr_ui(unsafe_mpz q, unsafe_mpz r, const unsafe_mpz n, long unsigned int d);
func (n Mpz) FDivQRUi(d uint) (q Mpz, r Mpz, ui uint, err error) {
	q, r = MpzInit(), MpzInit()
	qPtr, rPtr, nPtr, err := get3Pointers(q, r, n)
	if err != nil {
		return q, r, 0, err
	}
	uiLong := C.pmpz_fdiv_qr_ui(qPtr, rPtr, nPtr, C.ulong(d))
	return q, r, uint(uiLong), nil
}

//unsigned long int pmpz_fdiv_r_ui(const unsafe_mpz r, const unsafe_mpz n, long unsigned int d);
func (m Mpz) FDivRUi(d uint) (r Mpz, ui uint, err error) {
	r = MpzInit()
	rPtr, mPtr, err := get2Pointers(r, m)
	if err != nil {
		return r, 0, err
	}
	uiLong := C.pmpz_fdiv_r_ui(rPtr, mPtr, C.ulong(d))
	return r, uint(uiLong), err
}

//unsigned long int pmpz_fdiv_ui(const unsafe_mpz n, unsigned long int d);
func (m Mpz) FDivUi(d uint) (uint, error) {
	mPtr, err := m.Ptr()
	if err != nil {
		return 0, err
	}
	ui := C.pmpz_fdiv_ui(mPtr, C.ulong(d))
	return uint(ui), err
}

//unsigned long int pmpz_mod_ui(unsafe_mpz r, const unsafe_mpz n, unsigned long int d);
func (m Mpz) ModUi(d uint) (r Mpz, ui uint, err error) {
	r = MpzInit()
	rPtr, mPtr, err := get2Pointers(r, m)
	if err != nil {
		return r, 0, err
	}
	uiLong := C.pmpz_mod_ui(rPtr, mPtr, C.ulong(d))
	return r, uint(uiLong), err
}

//unsigned long int pmpz_tdiv_q_ui(unsafe_mpz q, const unsafe_mpz n, unsigned long int d);
func (m Mpz) TDivQUi(d uint) (q Mpz, ui uint, err error) {
	q = MpzInit()
	qPtr, mPtr, err := get2Pointers(q, m)
	if err != nil {
		return q, 0, err
	}
	uiLong := C.pmpz_tdiv_q_ui(qPtr, mPtr, C.ulong(d))
	return q, uint(uiLong), nil
}

//unsigned long int pmpz_tdiv_qr_ui(unsafe_mpz q, unsafe_mpz r, const unsafe_mpz n, unsigned long int d);.
func (m Mpz) TDivQRUi(d uint) (q Mpz, r Mpz, ui uint) {
	q, r = MpzInit(), MpzInit()
	qPtr, rPtr, mPtr, err := get3Pointers(q, r, m)
	if err != nil {
		return q, r, 0, err
	}
	uiLong := C.pmpz_tdiv_qr_ui(qPtr, rPtr, mPtr, C.ulong(d))
	return q, r, uint(uiLong), nil
}

//unsigned long int pmpz_tdiv_r_ui(unsafe_mpz r, const unsafe_mpz n, unsigned long int d);
func (m Mpz) TDivRUi(d uint) (r Mpz, ui uint, err error) {
	r = MpzInit()
	rPtr, mPtr, err := get2Pointers(r, m)
	if err != nil {
		return r, 0, err
	}
	uiLong := C.pmpz_tdiv_r_ui(rPtr, mPtr, C.ulong(d))
	return r, uint(uiLong), nil
}

//unsigned long int pmpz_tdiv_ui(const unsafe_mpz n, unsigned long int d);
func (m Mpz) TDivUi(d uint) (uint, error) {
	mPtr, err := m.Ptr()
	if err != nil {
		return 0, err
	}
	ret := C.pmpz_tdiv_ui(mPtr, C.ulong(d))
	return uint(ret), nil
}

//void pmpz_cdiv_qr(unsafe_mpz q, unsafe_mpz r, const unsafe_mpz n, const unsafe_mpz d);
func (m Mpz) CDivQR(d Mpz) (q Mpz, r Mpz, err error) {
	qPtr, rPtr, mPtr, dPtr, err := get4Pointers(q, r, m, d)
	if err != nil {
		return q, r, err
	}
	C.pmpz_cdiv_qr(qPtr, rPtr, mPtr, dPtr)
	return q, r, nil
}

//void pmpz_fdiv_qr(unsafe_mpz q, unsafe_mpz r, const unsafe_mpz n, const unsafe_mpz d);
func (m Mpz) FDivQR(d Mpz) (q Mpz, r Mpz, err error) {
	qPtr, rPtr, mPtr, dPtr, err := get4Pointers(q, r, m, d)
	if err != nil {
		return q, r, err
	}
	C.pmpz_fdiv_qr(qPtr, rPtr, mPtr, dPtr)
	return q, r, nil
}

//void pmpz_tdiv_qr(unsafe_mpz q, unsafe_mpz r, const unsafe_mpz n, const unsafe_mpz d);
func (m Mpz) TDivQR(d Mpz) (q Mpz, r Mpz, err error) {
	qPtr, rPtr, mPtr, dPtr, err := get4Pointers(q, r, m, d)
	if err != nil {
		return q, r, err
	}
	C.pmpz_tdiv_qr(qPtr, rPtr, mPtr, dPtr)
	return q, r, nil
}

//unsafe_mpz pmpz_cdiv_q(const unsafe_mpz n, const unsafe_mpz d);
func (m Mpz) CDivQ(d Mpz) (Mpz, error) {
	mPtr, dPtr, err := get2Pointers(m, d)
	if err != nil {
		return MpzInit(), err
	}
	ptr := C.pmpz_cdiv_q(mPtr, dPtr)
	return mpzFromPtr(ptr)
}

//unsafe_mpz pmpz_cdiv_q_2exp(const unsafe_mpz n, const mp_bitcnt_t b);
func (m Mpz) CDivQ2Exp(bitcnt uint) (Mpz, error) {
	mPtr, err := m.Ptr()
	if err != nil {
		return MpzInit(), err
	}
	ptr := C.pmpz_cdiv_q_2exp(mPtr, C.ulong(bitcnt))
	return mpzFromPtr(ptr)
}

//unsafe_mpz pmpz_cdiv_r(const unsafe_mpz n, const unsafe_mpz d);
func (m Mpz) CDivR(d Mpz) (Mpz, error) {
	mPtr, dPtr, err := get2Pointers(m, d)
	if err != nil {
		return MpzInit(), err
	}
	ptr := C.pmpz_cdiv_r(mPtr, dPtr)
	return mpzFromPtr(ptr)
}

//unsafe_mpz pmpz_cdiv_r_2exp(const unsafe_mpz n, const mp_bitcnt_t b);
func (m Mpz) CDivR2Exp(bitcnt uint) (Mpz, error) {
	mPtr, err := m.Ptr()
	if err != nil {
		return MpzInit(), err
	}
	ptr := C.pmpz_cdiv_r_2exp(mPtr, C.ulong(bitcnt))
	return mpzFromPtr(ptr)
}

//unsafe_mpz pmpz_divexact(const unsafe_mpz n, const unsafe_mpz d);
func (m Mpz) CDivExact(d Mpz) (Mpz, error) {
	mPtr, dPtr, err := get2Pointers(m, d)
	if err != nil {
		return MpzInit(), err
	}
	ptr := C.pmpz_divexact(mPtr, dPtr)
	return mpzFromPtr(ptr)
}

//unsafe_mpz pmpz_divexact_ui(const unsafe_mpz n, unsigned long d);
func (m Mpz) CDivExactUi(d C.ulong) (Mpz, error) {
	mPtr, err := m.Ptr()
	if err != nil {
		return MpzInit(), err
	}
	ptr := C.pmpz_divexact_ui(mPtr, d)
	return mpzFromPtr(ptr)
}

//unsafe_mpz pmpz_fdiv_q(const unsafe_mpz n, const unsafe_mpz d);
func (m Mpz) FDivQ(d Mpz) (Mpz, error) {
	mPtr, dPtr, err := get2Pointers(m, d)
	if err != nil {
		return MpzInit(), err
	}
	ptr := C.pmpz_fdiv_q(mPtr, dPtr)
	return mpzFromPtr(ptr)
}

//unsafe_mpz pmpz_fdiv_q_2exp(const unsafe_mpz n, const mp_bitcnt_t b);
func (m Mpz) FDivQ2Exp(bitcnt C.ulong) (Mpz, error) {
	mPtr, err := m.Ptr()
	if err != nil {
		return MpzInit(), err
	}
	ptr := C.pmpz_fdiv_q_2exp(mPtr, bitcnt)
	return mpzFromPtr(ptr)
}

//unsafe_mpz pmpz_fdiv_r(const unsafe_mpz n, const unsafe_mpz d);
func (m Mpz) FDivR(d Mpz) (Mpz, error) {
	mPtr, dPtr, err := get2Pointers(m, d)
	if err != nil {
		return MpzInit(), err
	}
	ptr := C.pmpz_fdiv_r(mPtr, dPtr)
	return mpzFromPtr(ptr)
}

//unsafe_mpz pmpz_fdiv_r_2exp(const unsafe_mpz n, const mp_bitcnt_t b);
func (m Mpz) FDivR2Exp(bitcnt C.ulong) (Mpz, error) {

	mPtr, err := m.Ptr()
	if err != nil {
		return MpzInit(), err
	}
	ptr := C.pmpz_fdiv_r_2exp(mPtr, bitcnt)
	return mpzFromPtr(ptr)
}

//unsafe_mpz pmpz_mod(const unsafe_mpz n, const unsafe_mpz d);
func (m Mpz) Mod(bitcnt C.ulong) (Mpz, error) {
	mPtr, err := m.Ptr()
	if err != nil {
		return MpzInit(), err
	}
	ptr := C.pmpz_fdiv_r_2exp(mPtr, bitcnt)
	return mpzFromPtr(ptr)
}

//unsafe_mpz pmpz_tdiv_q(const unsafe_mpz n, const unsafe_mpz d);
func (m Mpz) TDivQ(d Mpz) (Mpz, error) {

	mPtr, dPtr, err := get2Pointers(m, d)
	if err != nil {
		return MpzInit(), err
	}
	ptr := C.pmpz_tdiv_q(mPtr, dPtr)
	return mpzFromPtr(ptr)
}

//unsafe_mpz pmpz_tdiv_q_2exp(const unsafe_mpz n, const mp_bitcnt_t b);
func (m Mpz) TDivQ2Exp(bitcnt C.ulong) (Mpz, error) {
	mPtr, err := m.Ptr()
	if err != nil {
		return MpzInit(), err
	}
	ptr := C.pmpz_tdiv_q_2exp(mPtr, bitcnt)
	return mpzFromPtr(ptr)
}

//unsafe_mpz pmpz_tdiv_r(const unsafe_mpz n, const unsafe_mpz d);
func (m Mpz) TDivR(d Mpz) (Mpz, error) {
	mPtr, dPtr, err := get2Pointers(m, d)
	if err != nil {
		return MpzInit(), err
	}
	ptr := C.pmpz_tdiv_r(mPtr, dPtr)
	return mpzFromPtr(ptr)
}

//unsafe_mpz pmpz_tdiv_r_2exp(const unsafe_mpz n, const mp_bitcnt_t b);
func (m Mpz) TDivR2Exp(bitcnt C.ulong) (Mpz, error) {
	mPtr, err := m.Ptr()
	if err != nil {
		return MpzInit(), err
	}
	ptr := C.pmpz_tdiv_r_2exp(mPtr, bitcnt)
	return mpzFromPtr(ptr)
}
