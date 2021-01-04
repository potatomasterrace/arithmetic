package mpz

// #cgo LDFLAGS: -lgmp
// #include "mpz.h"
import "C"

// OK
// Number Theoretic Functions
// int pmpz_invert(unsafe_mpz rop, const unsafe_mpz op1, const unsafe_mpz op2);

func (m Mpz) MpzInvert(op Mpz) (Mpz, C.int) {
	rop := MpzInit()
	r := C.pmpz_invert(rop.Ptr(), m.Ptr(), op.Ptr())
	return rop, r
}

// int pmpz_jacobi(const unsafe_mpz a, const unsafe_mpz b);
func (m Mpz) MpzJacobi(op Mpz) C.int {
	return C.pmpz_jacobi(m.Ptr(), op.Ptr())
}

// int pmpz_kronecker(const unsafe_mpz a, const unsafe_mpz b);
func (m Mpz) MpzKronecker(op Mpz) C.int {
	return C.pmpz_kronecker(m.Ptr(), op.Ptr())
}

// int pmpz_kronecker_si(const unsafe_mpz a, long b);
func (m Mpz) MpzKroneckerSi(op C.long) C.int {
	return C.pmpz_kronecker_si(m.Ptr(), op)
}

// int pmpz_kronecker_ui(const unsafe_mpz a, unsigned long b);
func (m Mpz) MpzKroneckerUi(op C.long) C.int {
	return C.pmpz_kronecker_si(m.Ptr(), op)
}

// int pmpz_legendre(const unsafe_mpz a, const unsafe_mpz p);
func (m Mpz) MpzLegendre(p Mpz) C.int {
	return C.pmpz_legendre(m.Ptr(), p.Ptr())
}

// int pmpz_probab_prime_p(const unsafe_mpz n, int reps);
func (m Mpz) MpzProbabPrimeP(reps C.int) C.int {
	return C.pmpz_probab_prime_p(m.Ptr(), reps)
}

// int pmpz_si_kronecker(long a, const unsafe_mpz b);
func (m Mpz) MpzSiKronecker(reps C.int) C.int {
	return C.pmpz_probab_prime_p(m.Ptr(), reps)
}

// int pmpz_ui_kronecker(unsigned long a, const unsafe_mpz b);
func UiKroneckerMpz(a C.ulong, b Mpz) C.int {
	return C.pmpz_ui_kronecker(a, b.Ptr())
}

// mp_bitcnt_t pmpz_remove(unsafe_mpz rop, const unsafe_mpz op, const mpz_t f);
func (m Mpz) MpzRemove(op Mpz, f Mpz) C.ulong {
	rop := MpzInit()
	return C.pmpz_remove(rop.Ptr(), op.Ptr(), f.Ptr())
}

// unsafe_mpz pmpz_bin_ui(const unsafe_mpz n, unsigned long int k);
func (m Mpz) MpzBinUi(n C.ulong) Mpz {
	ptr := C.pmpz_2fac_ui(n)
	return mpzFromPtr(ptr)
}

// unsafe_mpz pmpz_bin_uiui(unsigned long int n, unsigned long int k);
func (m Mpz) MpzBinUiUi(n C.ulong, k C.ulong) Mpz {
	ptr := C.pmpz_2fac_ui(n)
	return mpzFromPtr(ptr)
}

// unsafe_mpz pmpz_gcd(const unsafe_mpz op1, const unsafe_mpz op2);
func (op1 Mpz) MpzGcd(op2 Mpz) Mpz {
	ptr := C.pmpz_gcd(op1.Ptr(), op2.Ptr())
	return mpzFromPtr(ptr)
}

// unsafe_mpz pmpz_lcm(const unsafe_mpz op1, const unsafe_mpz op2);
func (op1 Mpz) MpzLcm(op2 Mpz) Mpz {
	ptr := C.pmpz_lcm(op1.Ptr(), op2.Ptr())
	return mpzFromPtr(ptr)
}

// unsafe_mpz pmpz_lcm_ui(const unsafe_mpz op1, unsigned long op2);
func (op1 Mpz) MpzLcmUi(op2 C.ulong) Mpz {
	ptr := C.pmpz_lcm_ui(op1.Ptr(), op2)
	return mpzFromPtr(ptr)
}

// unsafe_mpz pmpz_nextprime(const unsafe_mpz op);
func (op Mpz) MpzNextPrime() Mpz {
	ptr := C.pmpz_nextprime(op.Ptr())
	return mpzFromPtr(ptr)
}

// unsigned long int pmpz_gcd_ui(unsafe_mpz rop, const unsafe_mpz op1, unsigned long int op2);
func (op1 Mpz) MpzGcdUi(op2 C.ulong) Mpz {
	rop := MpzInit()
	C.pmpz_gcd_ui(rop.Ptr(), op1.Ptr(), op2)
	return rop
}

// void pmpz_fib2_ui(unsafe_mpz fn, unsafe_mpz fnsub1, unsigned long int n);
func MpzFib2Ui(n C.ulong) (fn Mpz, fnSub1 Mpz) {
	fn, fnSub1 = MpzInit(), MpzInit()
	C.pmpz_fib2_ui(fn.Ptr(), fnSub1.Ptr(), n)
	return fn, fnSub1
}

// void pmpz_gcdext(unsafe_mpz g, unsafe_mpz s, unsafe_mpz t, const unsafe_mpz a, const unsafe_mpz b);
func MpzGcdext(a Mpz, b Mpz) (g Mpz, s Mpz, t Mpz) {
	g, s, t = MpzInit(), MpzInit(), MpzInit()
	C.pmpz_gcdext(g.Ptr(), s.Ptr(), t.Ptr(), a.Ptr(), b.Ptr())
	return g, s, t
}

// void pmpz_lucnum2_ui(unsafe_mpz ln, unsafe_mpz lnsub1, unsigned long int n);
func MpzLucNum2Ui(n C.ulong) (ln Mpz, lnsub1 Mpz) {
	ln, lnsub1 = MpzInit(), MpzInit()
	C.pmpz_lucnum2_ui(ln.Ptr(), lnsub1.Ptr(), n)
	return ln, lnsub1
}
