#include "mpz.h"

// Initialisation and assignement
// https://gmplib.org/manual/Assigning-Integers#Assigning-Integers
// Proxies to mpz_set, mpz_set_ui, mpz_set_si, mpz_set_d, mpz_set_q, mpz_set_f and mpz_set_str
unsafe_mpz pmpz_init()
{
    p_mpz_init(ptr);
    return (unsafe_mpz)ptr;
}

unsafe_mpz pmpz_set_mpz(const unsafe_mpz op)
{
    p_mpz mpz_op = (p_mpz)op;
    wrap_mpz_function(mpz_set, *mpz_op);
}

unsafe_mpz pmpz_set_ui(const unsigned long int op)
{
    wrap_mpz_function(mpz_set_ui, op);
}

unsafe_mpz pmpz_set_si(const signed long int op)
{
    wrap_mpz_function(mpz_set_si, op);
}

unsafe_mpz pmpz_set_double(const double op)
{
    wrap_mpz_function(mpz_set_d, op);
}

unsafe_mpz pmpz_set_quotient(const mpq_t op)
{
    wrap_mpz_function(mpz_set_q, op);
}

unsafe_mpz pmpz_set_float(const mpf_t op)
{
    wrap_mpz_function(mpz_set_f, op);
}

unsafe_mpz pmpz_set_str(const char *s, const int base)
{
    p_mpz_init(rop);
    if (mpz_set_str(*rop, s, 10) != 0)
    {
        return (unsafe_mpz)0;
    }
    return rop;
}
unsafe_mpz pmpz_ui_pow_ui(unsigned long int base, unsigned long int exp)
{
    wrap_mpz_function(mpz_ui_pow_ui, base, exp);
}

unsafe_mpz pmpz_fac_ui(unsigned long int n)
{
    wrap_mpz_function(mpz_fac_ui, n);
}

unsafe_mpz pmpz_2fac_ui(unsigned long int n)
{
    wrap_mpz_function(mpz_2fac_ui, n);
}

unsafe_mpz pmpz_mfac_uiui(unsigned long int n, unsigned long int m)
{
    wrap_mpz_function(mpz_mfac_uiui, n, m);
}

unsafe_mpz pmpz_primorial_ui(unsigned long int n)
{
    wrap_mpz_function(mpz_primorial_ui, n);
}

unsafe_mpz pmpz_bin_ui(const unsafe_mpz n, unsigned long int k)
{
    wrap_mpz_function(mpz_bin_ui, n, k);
}

unsafe_mpz pmpz_bin_uiui(unsigned long int n, unsigned long int k)
{
    wrap_mpz_function(mpz_bin_uiui, n, k);
}

unsafe_mpz pmpz_fib_ui(unsigned long int n)
{
    wrap_mpz_function(mpz_fib_ui, n);
}


void clear_mpz(const unsafe_mpz value)
{
    p_mpz num = (p_mpz)value;
    mpz_clear(*num);
    FREE_PROXY("mpz",num);
}


// Conversion Functions
// https://gmplib.org/manual/Converting-Integers#Converting-Integers

unsigned long int pmpz_get_ui(const unsafe_mpz op)
{
    ref_to_pmpz(op);
    return mpz_get_ui(ref(op));
}

signed long int pmpz_get_si(const unsafe_mpz op)
{
    ref_to_pmpz(op);
    return mpz_get_si(ref(op));
}

double pmpz_get_d(const unsafe_mpz op)
{
    ref_to_pmpz(op);
    return mpz_get_d(ref(op));
}

double pmpz_get_d_2exp(signed long int *exp, const unsafe_mpz op)
{
    ref_to_pmpz(op);
    return mpz_get_d_2exp(exp, ref(op));
}

char *pmpz_get_str(const unsafe_mpz n, const int base)
{
    ref_to_pmpz(n);
    int str_len = mpz_sizeinbase(ref(n), base) + 1;
    char *str = (char *)malloc(str_len);
    mpz_get_str(str, base, ref(n));
    return str;
}

// Arithmetic Functions
// https://gmplib.org/manual/Integer-Arithmetic#Integer-Arithmetic

unsafe_mpz pmpz_add(const unsafe_mpz op1, const unsafe_mpz op2)
{
    ref_to_pmpz(op1);
    ref_to_pmpz(op2);
    wrap_mpz_function(mpz_add, ref(op1), ref(op2));
}

unsafe_mpz pmpz_add_ui(const unsafe_mpz op1, const unsigned long int op2)
{
    ref_to_pmpz(op1);
    wrap_mpz_function(mpz_add_ui, ref(op1), op2);
}

unsafe_mpz pmpz_sub(const unsafe_mpz op1, const unsafe_mpz op2)
{
    ref_to_pmpz(op1);
    ref_to_pmpz(op2);
    wrap_mpz_function(mpz_sub, ref(op1), ref(op2));
}

unsafe_mpz pmpz_sub_ui(const unsafe_mpz op1, const unsigned long int op2)
{
    ref_to_pmpz(op1);
    wrap_mpz_function(mpz_sub_ui, ref(op1), op2);
}

unsafe_mpz pmpz_ui_sub(const unsigned long int op1, const unsafe_mpz op2)
{
    ref_to_pmpz(op2);
    wrap_mpz_function(mpz_ui_sub, op1, ref(op2));
}

unsafe_mpz pmpz_mul(const unsafe_mpz op1, const unsafe_mpz op2)
{
    ref_to_pmpz(op1);
    ref_to_pmpz(op2);
    wrap_mpz_function(mpz_mul, ref(op1), ref(op2));
}

unsafe_mpz pmpz_mul_si(const unsafe_mpz op1, long int op2)
{
    ref_to_pmpz(op1);
    wrap_mpz_function(mpz_mul_si, ref(op1), op2);
}

unsafe_mpz pmpz_mul_ui(const unsafe_mpz op1, unsigned long int op2)
{
    ref_to_pmpz(op1);
    wrap_mpz_function(mpz_mul_ui, ref(op1), op2);
}

void pmpz_unsafe_addmul(unsafe_mpz rop, const unsafe_mpz op1, const unsafe_mpz op2)
{
    ref_to_pmpz(rop);
    ref_to_pmpz(op1);
    ref_to_pmpz(op2);
    mpz_addmul(ref(rop), ref(op1), ref(op2));
}

void pmpz_unsafe_addmul_ui(unsafe_mpz rop, const unsafe_mpz op1, const unsigned long int op2)
{
    ref_to_pmpz(rop);
    ref_to_pmpz(op1);
    mpz_addmul_ui(ref(rop), ref(op1), op2);
}

void pmpz_unsafe_submul(unsafe_mpz rop, const unsafe_mpz op1, const unsafe_mpz op2)
{
    ref_to_pmpz(rop);
    ref_to_pmpz(op1);
    ref_to_pmpz(op2);
    mpz_submul(ref(rop), ref(op1), op2);
}

void pmpz_unsafe_submul_ui(unsafe_mpz rop, const unsafe_mpz op1, const unsigned long int op2)
{
    ref_to_pmpz(rop);
    ref_to_pmpz(op1);
    mpz_submul_ui(ref(rop), ref(op1), op2);
}

unsafe_mpz pmpz_mul_2exp(const unsafe_mpz op1, const mp_bitcnt_t op2)
{
    ref_to_pmpz(op1);
    wrap_mpz_function(mpz_mul_2exp, ref(op1), op2);
}

unsafe_mpz pmpz_neg(const unsafe_mpz op)
{
    ref_to_pmpz(op);
    wrap_mpz_function(mpz_neg, ref(op));
}

unsafe_mpz pmpz_abs(const unsafe_mpz op)
{
    ref_to_pmpz(op);
    wrap_mpz_function(mpz_abs, ref(op));
}

// Division Functions
// https://gmplib.org/manual/Integer-Division

unsafe_mpz pmpz_cdiv_q(const unsafe_mpz n, const unsafe_mpz d)
{
    ref_to_pmpz(n);
    ref_to_pmpz(d);
    wrap_mpz_function(mpz_cdiv_q, ref(n), ref(d));
}

unsafe_mpz pmpz_cdiv_r(const unsafe_mpz n, const unsafe_mpz d)
{
    ref_to_pmpz(n);
    ref_to_pmpz(d);
    wrap_mpz_function(mpz_cdiv_r, ref(n), ref(d));
}

void pmpz_cdiv_qr(unsafe_mpz q, unsafe_mpz r, const unsafe_mpz n, const unsafe_mpz d)
{
    ref_to_pmpz(r);
    ref_to_pmpz(q);
    ref_to_pmpz(n);
    ref_to_pmpz(d);
    mpz_cdiv_qr(ref(q), ref(r), ref(n), ref(d));
}

unsigned long int pmpz_cdiv_q_ui(unsafe_mpz q, const unsafe_mpz n, unsigned long int d)
{
    ref_to_pmpz(q);
    ref_to_pmpz(n);
    return mpz_cdiv_q_ui(ref(q), ref(n), d);
}

unsigned long int pmpz_cdiv_r_ui(unsafe_mpz r, const unsafe_mpz n, unsigned long int d)
{
    ref_to_pmpz(r);
    ref_to_pmpz(n);
    return mpz_cdiv_r_ui(ref(r), ref(n), d);
}

unsigned long int pmpz_cdiv_qr_ui(unsafe_mpz q, unsafe_mpz r, const unsafe_mpz n, unsigned long int d)
{
    ref_to_pmpz(r);
    ref_to_pmpz(q);
    ref_to_pmpz(n);
    ref_to_pmpz(d);
    return mpz_cdiv_qr_ui(ref(q), ref(r), ref(n), d);
}

unsigned long int pmpz_cdiv_ui(const unsafe_mpz n, unsigned long int d)
{
    ref_to_pmpz(n);
    return mpz_cdiv_ui(ref(n), d);
}

unsafe_mpz pmpz_cdiv_q_2exp(const unsafe_mpz n, const mp_bitcnt_t b)
{
    ref_to_pmpz(n);
    wrap_mpz_function(mpz_cdiv_q_2exp, ref(n), b);
}

unsafe_mpz pmpz_cdiv_r_2exp(const unsafe_mpz n, const mp_bitcnt_t b)
{
    ref_to_pmpz(n);
    wrap_mpz_function(mpz_cdiv_r_2exp, ref(n), b);
}

unsafe_mpz pmpz_fdiv_q(const unsafe_mpz n, const unsafe_mpz d)
{
    ref_to_pmpz(n);
    wrap_mpz_function(mpz_fdiv_q, ref(n), d);
}

unsafe_mpz pmpz_fdiv_r(const unsafe_mpz n, const unsafe_mpz d)
{
    ref_to_pmpz(n);
    ref_to_pmpz(d);
    wrap_mpz_function(mpz_fdiv_q, ref(n), ref(d));
}

void pmpz_fdiv_qr(unsafe_mpz q, unsafe_mpz r, const unsafe_mpz n, const unsafe_mpz d)
{
    ref_to_pmpz(q);
    ref_to_pmpz(r);
    ref_to_pmpz(n);
    ref_to_pmpz(d);
    mpz_fdiv_qr(ref(q), ref(r), ref(n), ref(d));
}

unsigned long int pmpz_fdiv_q_ui(const unsafe_mpz q, const unsafe_mpz n, long unsigned int d)
{
    ref_to_pmpz(n);
    return mpz_fdiv_q_ui(ref(n), ref(n), d);
}

unsigned long int pmpz_fdiv_r_ui(const unsafe_mpz r, const unsafe_mpz n, long unsigned int d)
{
    ref_to_pmpz(r);
    ref_to_pmpz(n);
    return mpz_fdiv_q_ui(ref(r), ref(n), d);
}

unsigned long int pmpz_fdiv_qr_ui(unsafe_mpz q, unsafe_mpz r, const unsafe_mpz n, long unsigned int d)
{
    ref_to_pmpz(q);
    ref_to_pmpz(r);
    ref_to_pmpz(n);
    return mpz_fdiv_qr_ui(ref(q), ref(r), ref(n), d);
}

unsigned long int pmpz_fdiv_ui(const unsafe_mpz n, unsigned long int d)
{
    ref_to_pmpz(n);
    return mpz_fdiv_ui(ref(n), d);
}

unsafe_mpz pmpz_fdiv_q_2exp(const unsafe_mpz n, const mp_bitcnt_t b)
{
    ref_to_pmpz(n);
    wrap_mpz_function(mpz_fdiv_q_2exp, ref(n), b);
}

unsafe_mpz pmpz_fdiv_r_2exp(const unsafe_mpz n, const mp_bitcnt_t b)
{
    ref_to_pmpz(n);
    wrap_mpz_function(mpz_fdiv_r_2exp, ref(n), b);
}

unsafe_mpz pmpz_tdiv_q(const unsafe_mpz n, const unsafe_mpz d)
{
    ref_to_pmpz(n);
    ref_to_pmpz(d);
    wrap_mpz_function(mpz_tdiv_q, ref(n), ref(d));
}

unsafe_mpz pmpz_tdiv_r(const unsafe_mpz n, const unsafe_mpz d)
{
    ref_to_pmpz(n);
    ref_to_pmpz(d);
    wrap_mpz_function(mpz_tdiv_q, ref(n), ref(d));
}

void pmpz_tdiv_qr(unsafe_mpz q, unsafe_mpz r, const unsafe_mpz n, const unsafe_mpz d)
{
    ref_to_pmpz(q);
    ref_to_pmpz(r);
    ref_to_pmpz(n);
    ref_to_pmpz(d);
    mpz_tdiv_qr(ref(q), ref(r), ref(n), ref(d));
}

unsigned long int pmpz_tdiv_q_ui(unsafe_mpz q, const unsafe_mpz n, unsigned long int d)
{
    ref_to_pmpz(q);
    ref_to_pmpz(n);
    return mpz_tdiv_q_ui(ref(q), ref(n), d);
}

unsigned long int pmpz_tdiv_r_ui(unsafe_mpz r, const unsafe_mpz n, unsigned long int d)
{
    ref_to_pmpz(r);
    ref_to_pmpz(n);
    return mpz_tdiv_q_ui(ref(r), ref(n), d);
}

unsigned long int pmpz_tdiv_qr_ui(unsafe_mpz q, unsafe_mpz r, const unsafe_mpz n, unsigned long int d)
{
    ref_to_pmpz(q);
    ref_to_pmpz(r);
    ref_to_pmpz(n);
    return mpz_tdiv_qr_ui(ref(q), ref(r), ref(n), d);
}

unsigned long int pmpz_tdiv_ui(const unsafe_mpz n, unsigned long int d)
{
    ref_to_pmpz(n);
    ref_to_pmpz(d);
    return mpz_tdiv_ui(ref(n), d);
}

unsafe_mpz pmpz_tdiv_q_2exp(const unsafe_mpz n, const mp_bitcnt_t b)
{
    ref_to_pmpz(n);
    wrap_mpz_function(mpz_tdiv_q_2exp, ref(n), b);
}

unsafe_mpz pmpz_tdiv_r_2exp(const unsafe_mpz n, const mp_bitcnt_t b)
{
    ref_to_pmpz(n);
    wrap_mpz_function(mpz_tdiv_r_2exp, ref(n), b);
}

unsafe_mpz pmpz_mod(const unsafe_mpz n, const unsafe_mpz d)
{
    ref_to_pmpz(n);
    ref_to_pmpz(d);
    wrap_mpz_function(mpz_mod, ref(n), ref(d));
}

unsigned long int pmpz_mod_ui(unsafe_mpz r, const unsafe_mpz n, unsigned long int d)
{
    ref_to_pmpz(r);
    ref_to_pmpz(n);
    return mpz_mod_ui(ref(r), ref(n), d);
}

unsafe_mpz pmpz_divexact(const unsafe_mpz n, const unsafe_mpz d)
{
    ref_to_pmpz(n);
    ref_to_pmpz(d);
    wrap_mpz_function(mpz_divexact, ref(n), ref(d));
}

unsafe_mpz pmpz_divexact_ui(const unsafe_mpz n, unsigned long d)
{
    ref_to_pmpz(n);
    wrap_mpz_function(mpz_divexact_ui, ref(n), d);
}

int pmpz_divisible_p(const unsafe_mpz n, const unsafe_mpz d)
{
    ref_to_pmpz(n);
    ref_to_pmpz(d);
    return mpz_divisible_p(ref(n), ref(d));
}

int pmpz_divisible_ui_p(const unsafe_mpz n, unsigned long int d)
{
    ref_to_pmpz(n);
    return mpz_divisible_ui_p(ref(n), d);
}

int pmpz_divisible_2exp_p(const unsafe_mpz n, mp_bitcnt_t b)
{
    ref_to_pmpz(n);
    return mpz_divisible_2exp_p(ref(n), b);
}

int pmpz_congruent_p(const unsafe_mpz n, const unsafe_mpz c, const unsafe_mpz d)
{
    ref_to_pmpz(n);
    ref_to_pmpz(c);
    ref_to_pmpz(d);
    return mpz_congruent_p(ref(n), ref(c), ref(d));
}

int pmpz_congruent_ui_p(const unsafe_mpz n, unsigned long int c, unsigned long int d)
{
    ref_to_pmpz(n);
    return mpz_congruent_ui_p(ref(n), c, d);
}

int pmpz_congruent_2exp_p(const unsafe_mpz n, const unsafe_mpz c, mp_bitcnt_t b)
{
    ref_to_pmpz(n);
    ref_to_pmpz(c);
    return mpz_congruent_2exp_p(ref(n), ref(c), b);
}

// Exponentiation Functions
// https://gmplib.org/manual/Integer-Exponentiation

unsafe_mpz pmpz_powm(const unsafe_mpz base, const unsafe_mpz exp, const unsafe_mpz mod)
{
    ref_to_pmpz(base);
    ref_to_pmpz(exp);
    ref_to_pmpz(mod);
    wrap_mpz_function(mpz_powm, ref(base), ref(exp), ref(mod));
}

unsafe_mpz pmpz_powm_ui(const unsafe_mpz base, unsigned long int exp, const unsafe_mpz mod)
{
    ref_to_pmpz(base);
    ref_to_pmpz(mod);
    wrap_mpz_function(mpz_powm_ui, ref(base), exp, ref(mod));
}

unsafe_mpz pmpz_powm_sec(const unsafe_mpz base, const unsafe_mpz exp, const unsafe_mpz mod)
{
    ref_to_pmpz(base);
    ref_to_pmpz(exp);
    ref_to_pmpz(mod);
    wrap_mpz_function(mpz_powm_sec, ref(base), exp, ref(mod));
}

unsafe_mpz pmpz_pow_ui(const unsafe_mpz base, unsigned long int exp)
{
    ref_to_pmpz(base);
    wrap_mpz_function(mpz_pow_ui, ref(base), exp);
}

// Root Extraction Functions
// https://gmplib.org/manual/Integer-Roots#Integer-Roots

int pmpz_root(unsafe_mpz rop, const unsafe_mpz op, unsigned long int n)
{
    ref_to_pmpz(rop);
    ref_to_pmpz(op);
    return pmpz_congruent_2exp_p(ref(rop), ref(op), n);
}

void pmpz_rootrem(unsafe_mpz root, unsafe_mpz rem, const unsafe_mpz u, unsigned long int n)
{
    ref_to_pmpz(root);
    ref_to_pmpz(rem);
    ref_to_pmpz(u);
    return mpz_rootrem(ref(root), ref(rem), ref(u), n);
}

unsafe_mpz pmpz_sqrt(const unsafe_mpz op)
{
    ref_to_pmpz(op);
    wrap_mpz_function(mpz_sqrt, ref(op));
}

void pmpz_sqrtrem(unsafe_mpz rop1, unsafe_mpz rop2, const unsafe_mpz op)
{
    ref_to_pmpz(rop1);
    ref_to_pmpz(rop2);
    ref_to_pmpz(op);
    mpz_sqrtrem(ref(rop1), ref(rop2), ref(op));
}

int pmpz_perfect_power_p(const unsafe_mpz op)
{
    ref_to_pmpz(op);
    return mpz_perfect_power_p(ref(op));
}

int pmpz_perfect_square_p(const unsafe_mpz op)
{
    ref_to_pmpz(op);
    return mpz_perfect_square_p(ref(op));
}

// Number Theoretic Functions
// https://gmplib.org/manual/Number-Theoretic-Functions

int pmpz_probab_prime_p(const unsafe_mpz n, int reps)
{
    ref_to_pmpz(n);
    return mpz_probab_prime_p(ref(n), reps);
}

unsafe_mpz pmpz_nextprime(const unsafe_mpz op)
{
    ref_to_pmpz(op);
    wrap_mpz_function(mpz_nextprime, ref(op));
}

unsafe_mpz pmpz_gcd(const unsafe_mpz op1, const unsafe_mpz op2)
{
    ref_to_pmpz(op1);
    ref_to_pmpz(op2);
    wrap_mpz_function(mpz_gcd, ref(op1), ref(op2));
}

unsigned long int pmpz_gcd_ui(unsafe_mpz rop, const unsafe_mpz op1, unsigned long int op2)
{
    ref_to_pmpz(rop);
    ref_to_pmpz(op1);
    return mpz_gcd_ui(ref(rop), ref(op1), op2);
}

void pmpz_gcdext(unsafe_mpz g, unsafe_mpz s, unsafe_mpz t, const unsafe_mpz a, const unsafe_mpz b)
{
    ref_to_pmpz(g);
    ref_to_pmpz(s);
    ref_to_pmpz(t);
    ref_to_pmpz(a);
    ref_to_pmpz(b);
    pmpz_gcdext(ref(g), ref(s), ref(t), ref(a), ref(b));
}

unsafe_mpz pmpz_lcm(const unsafe_mpz op1, const unsafe_mpz op2)
{
    ref_to_pmpz(op1);
    ref_to_pmpz(op2);
    wrap_mpz_function(mpz_lcm, ref(op1), ref(op2));
}

unsafe_mpz pmpz_lcm_ui(const unsafe_mpz op1, unsigned long op2)
{
    ref_to_pmpz(op1);
    wrap_mpz_function(mpz_lcm_ui, ref(op1), op2);
}

int pmpz_invert(unsafe_mpz rop, const unsafe_mpz op1, const unsafe_mpz op2)
{
    ref_to_pmpz(rop);
    ref_to_pmpz(op1);
    ref_to_pmpz(op2);
    return pmpz_invert(ref(rop), ref(op1), ref(op2));
}

int pmpz_jacobi(const unsafe_mpz a, const unsafe_mpz b)
{
    ref_to_pmpz(a);
    ref_to_pmpz(b);
    return mpz_jacobi(ref(a), ref(b));
}

int pmpz_legendre(const unsafe_mpz a, const unsafe_mpz p)
{
    ref_to_pmpz(a);
    ref_to_pmpz(p);
    return mpz_legendre(ref(a), ref(p));
}

int pmpz_kronecker(const unsafe_mpz a, const unsafe_mpz b)
{
    ref_to_pmpz(a);
    ref_to_pmpz(b);
    return mpz_legendre(ref(a), ref(b));
}

int pmpz_kronecker_si(const unsafe_mpz a, long b)
{
    ref_to_pmpz(a);
    return mpz_kronecker_si(ref(a), b);
}

int pmpz_kronecker_ui(const unsafe_mpz a, unsigned long b)
{
    ref_to_pmpz(a);
    return mpz_kronecker_ui(ref(a), b);
}

int pmpz_si_kronecker(long a, const unsafe_mpz b)
{
    ref_to_pmpz(b);
    return mpz_si_kronecker(a, ref(b));
}

int pmpz_ui_kronecker(unsigned long a, const unsafe_mpz b)
{
    ref_to_pmpz(b);
    return pmpz_ui_kronecker(a, ref(b));
}

mp_bitcnt_t pmpz_remove(unsafe_mpz rop,const unsafe_mpz op, const unsafe_mpz f)
{
    ref_to_pmpz(rop);
    ref_to_pmpz(op);
    ref_to_pmpz(f);
    return mpz_remove(ref(rop), ref(op), ref(f));
}


void pmpz_fib2_ui(unsafe_mpz fn, unsafe_mpz fnsub1, unsigned long int n)
{
    ref_to_pmpz(fn);
    ref_to_pmpz(fnsub1);
    pmpz_fib2_ui(ref(fn), ref(fnsub1), n);
}

unsafe_mpz pmpz_lucnum_ui(unsigned long int n)
{
    wrap_mpz_function(mpz_lucnum_ui, n);
}

void pmpz_lucnum2_ui(unsafe_mpz ln, unsafe_mpz lnsub1, unsigned long int n)
{
    ref_to_pmpz(ln);
    ref_to_pmpz(lnsub1);
    mpz_lucnum2_ui(ref(ln), ref(lnsub1), n);
}

// Comparison Functions
// https://gmplib.org/manual/Integer-Comparisons

int pmpz_cmp(const unsafe_mpz op1, const unsafe_mpz op2)
{
    ref_to_pmpz(op1);
    ref_to_pmpz(op2);
    return mpz_cmp(ref(op1), ref(op2));
}

int pmpz_cmp_d(const unsafe_mpz op1, double op2)
{
    ref_to_pmpz(op1);
    return mpz_cmp_d(ref(op1), op2);
}

int pmpz_cmp_si(const unsafe_mpz op1, signed long int op2)
{
    ref_to_pmpz(op1);
    return mpz_cmp_si(ref(op1), op2);
}

int pmpz_cmp_ui(const unsafe_mpz op1, unsigned long int op2)
{
    ref_to_pmpz(op1);
    return mpz_cmp_ui(ref(op1), op2);
}

int pmpz_cmpabs(const unsafe_mpz op1, const unsafe_mpz op2)
{
    ref_to_pmpz(op1);
    ref_to_pmpz(op2);
    return mpz_cmpabs(ref(op1), ref(op2));
}

int pmpz_cmpabs_d(const unsafe_mpz op1, double op2)
{
    ref_to_pmpz(op1);
    return mpz_cmp_ui(ref(op1), op2);
}

int pmpz_cmpabs_ui(const unsafe_mpz op1, unsigned long int op2)
{
    ref_to_pmpz(op1);
    return mpz_cmp_ui(ref(op1), op2);
}

int pmpz_sgn(const unsafe_mpz op)
{
    ref_to_pmpz(op);
    return mpz_sgn(ref(op));
}

// Logical and Bit Manipulation Functions
// https://gmplib.org/manual/Integer-Comparisons

unsafe_mpz pmpz_and(const unsafe_mpz op1, const unsafe_mpz op2)
{
    ref_to_pmpz(op1);
    ref_to_pmpz(op2);
    wrap_mpz_function(mpz_and, ref(op1), ref(op2));
}

unsafe_mpz pmpz_ior(const unsafe_mpz op1, const unsafe_mpz op2)
{
    ref_to_pmpz(op1);
    ref_to_pmpz(op2);
    wrap_mpz_function(mpz_ior, ref(op1), ref(op2));
}

unsafe_mpz pmpz_xor(const unsafe_mpz op1, const unsafe_mpz op2)
{
    ref_to_pmpz(op1);
    ref_to_pmpz(op2);
    wrap_mpz_function(mpz_xor, ref(op1), ref(op2));
}

unsafe_mpz pmpz_com(const unsafe_mpz op)
{
    ref_to_pmpz(op);
    wrap_mpz_function(mpz_com, ref(op));
}

mp_bitcnt_t pmpz_popcount(const unsafe_mpz op)
{
    ref_to_pmpz(op);
    return mpz_popcount(ref(op));
}

mp_bitcnt_t pmpz_hamdist(const unsafe_mpz op1, const unsafe_mpz op2)
{
    ref_to_pmpz(op1);
    ref_to_pmpz(op2);
    return mpz_hamdist(ref(op1), ref(op2));
}

mp_bitcnt_t pmpz_scan0(const unsafe_mpz op, mp_bitcnt_t starting_bit)
{
    ref_to_pmpz(op);
    return mpz_scan0(ref(op), starting_bit);
}

mp_bitcnt_t pmpz_scan1(const unsafe_mpz op, mp_bitcnt_t starting_bit)
{
    ref_to_pmpz(op);
    return mpz_scan1(ref(op), starting_bit);
}

void pmpz_setbit(const unsafe_mpz op, mp_bitcnt_t bit_index)
{
    ref_to_pmpz(op);
    mpz_setbit(ref(op), bit_index);
}

void pmpz_clrbit(unsafe_mpz rop, mp_bitcnt_t bit_index)
{
    ref_to_pmpz(rop);
    mpz_clrbit(ref(rop), bit_index);
}

void pmpz_combit(unsafe_mpz rop, mp_bitcnt_t bit_index)
{
    ref_to_pmpz(rop);
    mpz_clrbit(ref(rop), bit_index);
}

int pmpz_tstbit(const unsafe_mpz op, mp_bitcnt_t bit_index)
{
    ref_to_pmpz(op);
    return mpz_tstbit(ref(op), bit_index);
}

// Input and Output Functions
// https://gmplib.org/manual/I_002fO-of-Integers

size_t pmpz_out_str(FILE *stream, int base, const unsafe_mpz op)
{
    ref_to_pmpz(op);
    return mpz_out_str(stream, base, ref(op));
}

size_t pmpz_inp_str(unsafe_mpz rop, FILE *stream, int base)
{
    ref_to_pmpz(rop);
    return mpz_out_str(ref(rop), stream, base);
}

size_t pmpz_out_raw(FILE *stream, const unsafe_mpz op)
{
    ref_to_pmpz(op);
    return mpz_out_str(stream, ref(op));
}

size_t pmpz_inp_raw(unsafe_mpz rop, FILE *stream)
{
    ref_to_pmpz(rop);
    return mpz_out_str(ref(rop), stream);
}

// Random Number Functions
// https://gmplib.org/manual/Integer-Random-Numbers

unsafe_mpz pmpz_urandomb(gmp_randstate_t state, mp_bitcnt_t n)
{
    wrap_mpz_function(mpz_urandomb, state, n);
}

unsafe_mpz pmpz_urandomm(gmp_randstate_t state, const unsafe_mpz n)
{
    wrap_mpz_function(mpz_urandomm, state, n);
}

unsafe_mpz pmpz_rrandomb(gmp_randstate_t state, mp_bitcnt_t n)
{
    wrap_mpz_function(mpz_rrandomb, state, n);
}

unsafe_mpz pmpz_random(mp_size_t max_size)
{
    wrap_mpz_function(mpz_random, max_size);
}

unsafe_mpz pmpz_random2(mp_size_t max_size)
{
    wrap_mpz_function(mpz_random2, max_size);
}

// Integer Import and Export
// https://gmplib.org/manual/Integer-Import-and-Export

unsafe_mpz pmpz_import(size_t count, int order, size_t size, int endian, size_t nails, const void *op)
{
    wrap_mpz_function(mpz_import, count, order, size, endian, nails, op);
}

void *pmpz_export(int order, size_t size, int endian, size_t nails, const unsafe_mpz op)
{
    ref_to_pmpz(op);
    // int numb = 8 * size - nails;
    // size_t count = (mpz_sizeinbase(ref(op), 2) + numb - 1) / numb;
    return mpz_export(NULL, NULL, order, size, endian, nails, ref(op));
}

// Miscellaneous Functions
// https://gmplib.org/manual/Miscellaneous-Integer-Functions

int pmpz_fits_ulong_p(const unsafe_mpz op)
{
    ref_to_pmpz(op);
    return mpz_fits_ulong_p(ref(op));
}

int pmpz_fits_slong_p(const unsafe_mpz op)
{
    ref_to_pmpz(op);
    return mpz_fits_slong_p(ref(op));
}

int pmpz_fits_uint_p(const unsafe_mpz op)
{
    ref_to_pmpz(op);
    return mpz_fits_uint_p(ref(op));
}

int pmpz_fits_sint_p(const unsafe_mpz op)
{
    ref_to_pmpz(op);
    return mpz_fits_sint_p(ref(op));
}

int pmpz_fits_ushort_p(const unsafe_mpz op)
{
    ref_to_pmpz(op);
    return mpz_fits_ushort_p(ref(op));
}

int pmpz_fits_sshort_p(const unsafe_mpz op)
{
    ref_to_pmpz(op);
    return mpz_fits_sshort_p(ref(op));
}

int pmpz_even_p(const unsafe_mpz op)
{
    ref_to_pmpz(op);
    return mpz_even_p(ref(op));
}

size_t pmpz_sizeinbase(const unsafe_mpz op, int base)
{
    ref_to_pmpz(op);
    return mpz_sizeinbase(ref(op), base);
}

// Special Functions
// https://gmplib.org/manual/Integer-Special-Functions

size_t pmpz_size(const unsafe_mpz op)
{
    ref_to_pmpz(op);
    return mpz_size(ref(op));
}

// Commented out as these functions are hard to port
// mp_limb_t* mpz_limbs_read (const unsafe_mpz x){}
// void mpz_array_init (unsafe_mpz integer_array, mp_size_t array_size, mp_size_t fixed_num_bits){}
// void * _mpz_realloc (unsafe_mpz integer, mp_size_t new_alloc){}
// mp_limb_t mpz_getlimbn (const unsafe_mpz op, mp_size_t n){}
// void mpz_limbs_finish (unsafe_mpz x, mp_size_t s)
// mpz_srcptr mpz_roinit_n (unsafe_mpz x, const mp_limb_t *xp, mp_size_t xs)
// unsafe_mpz MPZ_ROINIT_N (mp_limb_t *xp, mp_size_t xs)
