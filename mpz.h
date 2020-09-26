#ifndef GO_MPZ_H
#define GO_MPZ_H

#include <gmp.h>
#include <malloc.h>

#define p_mpz mpz_t *

#define ref_to_pmpz(ref) \
    p_mpz p_##ref = (p_mpz)ref;

#define ref(value) \
    *p_##value

#define unsafe_mpz void *

// Macro for initializing mpz
#define p_mpz_init(value)                \
    /* allocate memory for reference */  \
    p_mpz value = malloc(sizeof(mpz_t)); \
    /* initialize mpz pointer  */        \
    mpz_init(*value);

// Macro for wrapping mpz functions
#define wrap_mpz_function(func, ...) \
    /* init return reference*/       \
    p_mpz_init(rop);                 \
    /* Call func */                  \
    func(*rop, __VA_ARGS__);         \
    /* return casted value */        \
    return (unsafe_mpz)rop;

void clear_mpz(const unsafe_mpz value);

// Initialisation and assignement
unsafe_mpz pmpz_init();
unsafe_mpz pmpz_set_double(const double op);
unsafe_mpz pmpz_set_float(const mpf_t op);
unsafe_mpz pmpz_set_mpz(const unsafe_mpz op);
unsafe_mpz pmpz_set_quotient(const mpq_t op);
unsafe_mpz pmpz_set_si(const signed long int op);
unsafe_mpz pmpz_set_str(const char *s, const int base);
unsafe_mpz pmpz_set_ui(const unsigned long int op);

// Conversion Functions
char *pmpz_get_str(const unsafe_mpz n, const int base);
double pmpz_get_d(const unsafe_mpz op);
double pmpz_get_d_2exp(signed long int *exp, const unsafe_mpz op);
signed long int pmpz_get_si(const unsafe_mpz op);
unsigned long int pmpz_get_ui(const unsafe_mpz op);

// Arithmetic Functions
unsafe_mpz pmpz_abs(const unsafe_mpz op);
unsafe_mpz pmpz_add(const unsafe_mpz op1, const unsafe_mpz op2);
unsafe_mpz pmpz_add_ui(const unsafe_mpz op1, const unsigned long int op2);
unsafe_mpz pmpz_mul(const unsafe_mpz op1, const unsafe_mpz op2);
unsafe_mpz pmpz_mul_2exp(const unsafe_mpz op1, const mp_bitcnt_t op2);
unsafe_mpz pmpz_mul_si(const unsafe_mpz op1, long int op2);
unsafe_mpz pmpz_mul_ui(const unsafe_mpz op1, unsigned long int op2);
unsafe_mpz pmpz_neg(const unsafe_mpz op);
unsafe_mpz pmpz_sub(const unsafe_mpz op1, const unsafe_mpz op2);
unsafe_mpz pmpz_sub_ui(const unsafe_mpz op1, const unsigned long int op2);
unsafe_mpz pmpz_ui_sub(const unsigned long int op1, const unsafe_mpz op2);
void pmpz_unsafe_addmul(unsafe_mpz rop, const unsafe_mpz op1, const unsafe_mpz op2);
void pmpz_unsafe_addmul_ui(unsafe_mpz rop, const unsafe_mpz op1, const unsigned long int op2);
void pmpz_unsafe_submul(unsafe_mpz rop, const unsafe_mpz op1, const unsafe_mpz op2);
void pmpz_unsafe_submul_ui(unsafe_mpz rop, const unsafe_mpz op1, const unsigned long int op2);

// Division Functions
int pmpz_congruent_2exp_p(const unsafe_mpz n, const unsafe_mpz c, mp_bitcnt_t b);
int pmpz_congruent_p(const unsafe_mpz n, const unsafe_mpz c, const unsafe_mpz d);
int pmpz_congruent_ui_p(const unsafe_mpz n, unsigned long int c, unsigned long int d);
int pmpz_divisible_2exp_p(const unsafe_mpz n, mp_bitcnt_t b);
int pmpz_divisible_p(const unsafe_mpz n, const unsafe_mpz d);
int pmpz_divisible_ui_p(const unsafe_mpz n, unsigned long int d);
unsafe_mpz pmpz_cdiv_q(const unsafe_mpz n, const unsafe_mpz d);
unsafe_mpz pmpz_cdiv_q_2exp(const unsafe_mpz n, const mp_bitcnt_t b);
unsafe_mpz pmpz_cdiv_r(const unsafe_mpz n, const unsafe_mpz d);
unsafe_mpz pmpz_cdiv_r_2exp(const unsafe_mpz n, const mp_bitcnt_t b);
unsafe_mpz pmpz_divexact(const unsafe_mpz n, const unsafe_mpz d);
unsafe_mpz pmpz_divexact_ui(const unsafe_mpz n, unsigned long d);
unsafe_mpz pmpz_fdiv_q(const unsafe_mpz n, const unsafe_mpz d);
unsafe_mpz pmpz_fdiv_q_2exp(const unsafe_mpz n, const mp_bitcnt_t b);
unsafe_mpz pmpz_fdiv_r(const unsafe_mpz n, const unsafe_mpz d);
unsafe_mpz pmpz_fdiv_r_2exp(const unsafe_mpz n, const mp_bitcnt_t b);
unsafe_mpz pmpz_mod(const unsafe_mpz n, const unsafe_mpz d);
unsafe_mpz pmpz_tdiv_q(const unsafe_mpz n, const unsafe_mpz d);
unsafe_mpz pmpz_tdiv_q_2exp(const unsafe_mpz n, const mp_bitcnt_t b);
unsafe_mpz pmpz_tdiv_r(const unsafe_mpz n, const unsafe_mpz d);
unsafe_mpz pmpz_tdiv_r_2exp(const unsafe_mpz n, const mp_bitcnt_t b);
unsigned long int pmpz_cdiv_q_ui(unsafe_mpz q, const unsafe_mpz n, unsigned long int d);
unsigned long int pmpz_cdiv_qr_ui(unsafe_mpz q, unsafe_mpz r, const unsafe_mpz n, unsigned long int d);
unsigned long int pmpz_cdiv_r_ui(unsafe_mpz r, const unsafe_mpz n, unsigned long int d);
unsigned long int pmpz_cdiv_ui(const unsafe_mpz n, unsigned long int d);
unsigned long int pmpz_fdiv_q_ui(const unsafe_mpz q, const unsafe_mpz n, long unsigned int d);
unsigned long int pmpz_fdiv_qr_ui(unsafe_mpz q, unsafe_mpz r, const unsafe_mpz n, long unsigned int d);
unsigned long int pmpz_fdiv_r_ui(const unsafe_mpz r, const unsafe_mpz n, long unsigned int d);
unsigned long int pmpz_fdiv_ui(const unsafe_mpz n, unsigned long int d);
unsigned long int pmpz_mod_ui(unsafe_mpz r, const unsafe_mpz n, unsigned long int d);
unsigned long int pmpz_tdiv_q_ui(unsafe_mpz q, const unsafe_mpz n, unsigned long int d);
unsigned long int pmpz_tdiv_qr_ui(unsafe_mpz q, unsafe_mpz r, const unsafe_mpz n, unsigned long int d);
unsigned long int pmpz_tdiv_r_ui(unsafe_mpz r, const unsafe_mpz n, unsigned long int d);
unsigned long int pmpz_tdiv_ui(const unsafe_mpz n, unsigned long int d);
void pmpz_cdiv_qr(unsafe_mpz q, unsafe_mpz r, const unsafe_mpz n, const unsafe_mpz d);
void pmpz_fdiv_qr(unsafe_mpz q, unsafe_mpz r, const unsafe_mpz n, const unsafe_mpz d);
void pmpz_tdiv_qr(unsafe_mpz q, unsafe_mpz r, const unsafe_mpz n, const unsafe_mpz d);

// Exponentiation Functions
unsafe_mpz pmpz_pow_ui(const unsafe_mpz base, unsigned long int exp);
unsafe_mpz pmpz_powm(const unsafe_mpz base, const unsafe_mpz exp, const unsafe_mpz mod);
unsafe_mpz pmpz_powm_sec(const unsafe_mpz base, const unsafe_mpz exp, const unsafe_mpz mod);
unsafe_mpz pmpz_powm_ui(const unsafe_mpz base, unsigned long int exp, const unsafe_mpz mod);
unsafe_mpz pmpz_ui_pow_ui(unsigned long int base, unsigned long int exp);

// Root Extraction Functions
int pmpz_perfect_power_p(const unsafe_mpz op);
int pmpz_perfect_square_p(const unsafe_mpz op);
int pmpz_root(unsafe_mpz rop, const unsafe_mpz op, unsigned long int n);
unsafe_mpz pmpz_sqrt(const unsafe_mpz op);
void pmpz_rootrem(unsafe_mpz root, unsafe_mpz rem, const unsafe_mpz u, unsigned long int n);
void pmpz_sqrtrem(unsafe_mpz rop1, unsafe_mpz rop2, const unsafe_mpz op);

// Number Theoretic Functions
int pmpz_invert(unsafe_mpz rop, const unsafe_mpz op1, const unsafe_mpz op2);
int pmpz_jacobi(const unsafe_mpz a, const unsafe_mpz b);
int pmpz_kronecker(const unsafe_mpz a, const unsafe_mpz b);
int pmpz_kronecker_si(const unsafe_mpz a, long b);
int pmpz_kronecker_ui(const unsafe_mpz a, unsigned long b);
int pmpz_legendre(const unsafe_mpz a, const unsafe_mpz p);
int pmpz_probab_prime_p(const unsafe_mpz n, int reps);
int pmpz_si_kronecker(long a, const unsafe_mpz b);
int pmpz_ui_kronecker(unsigned long a, const unsafe_mpz b);
mp_bitcnt_t pmpz_remove(unsafe_mpz rop, const unsafe_mpz op, const mpz_t f);
unsafe_mpz pmpz_2fac_ui(unsigned long int n);
unsafe_mpz pmpz_bin_ui(const unsafe_mpz n, unsigned long int k);
unsafe_mpz pmpz_bin_uiui(unsigned long int n, unsigned long int k);
unsafe_mpz pmpz_fac_ui(unsigned long int n);
unsafe_mpz pmpz_fib_ui(unsigned long int n);
unsafe_mpz pmpz_gcd(const unsafe_mpz op1, const unsafe_mpz op2);
unsafe_mpz pmpz_lcm(const unsafe_mpz op1, const unsafe_mpz op2);
unsafe_mpz pmpz_lcm_ui(const unsafe_mpz op1, unsigned long op2);
unsafe_mpz pmpz_lucnum_ui(unsigned long int n);
unsafe_mpz pmpz_mfac_uiui(unsigned long int n, unsigned long int m);
unsafe_mpz pmpz_nextprime(const unsafe_mpz op);
unsafe_mpz pmpz_primorial_ui(unsigned long int n);
unsigned long int pmpz_gcd_ui(unsafe_mpz rop, const unsafe_mpz op1, unsigned long int op2);
void pmpz_fib2_ui(unsafe_mpz fn, unsafe_mpz fnsub1, unsigned long int n);
void pmpz_gcdext(unsafe_mpz g, unsafe_mpz s, unsafe_mpz t, const unsafe_mpz a, const unsafe_mpz b);
void pmpz_lucnum2_ui(unsafe_mpz ln, unsafe_mpz lnsub1, unsigned long int n);

// Comparison Functions
int pmpz_cmp(const unsafe_mpz op1, const unsafe_mpz op2);
int pmpz_cmp_d(const unsafe_mpz op1, double op2);
int pmpz_cmp_si(const unsafe_mpz op1, signed long int op2);
int pmpz_cmp_ui(const unsafe_mpz op1, unsigned long int op2);
int pmpz_cmpabs(const unsafe_mpz op1, const unsafe_mpz op2);
int pmpz_cmpabs_d(const unsafe_mpz op1, double op2);
int pmpz_cmpabs_ui(const unsafe_mpz op1, unsigned long int op2);
int pmpz_sgn(const unsafe_mpz op);

// Logical and Bit Manipulation Functions
int pmpz_tstbit(const unsafe_mpz op, mp_bitcnt_t bit_index);
mp_bitcnt_t pmpz_hamdist(const unsafe_mpz op1, const unsafe_mpz op2);
mp_bitcnt_t pmpz_popcount(const unsafe_mpz op);
mp_bitcnt_t pmpz_scan0(const unsafe_mpz op, mp_bitcnt_t starting_bit);
mp_bitcnt_t pmpz_scan1(const unsafe_mpz op, mp_bitcnt_t starting_bit);
unsafe_mpz pmpz_and(const unsafe_mpz op1, const unsafe_mpz op2);
unsafe_mpz pmpz_com(const unsafe_mpz op);
unsafe_mpz pmpz_ior(const unsafe_mpz op1, const unsafe_mpz op2);
unsafe_mpz pmpz_xor(const unsafe_mpz op1, const unsafe_mpz op2);
void pmpz_clrbit(unsafe_mpz rop, mp_bitcnt_t bit_index);
void pmpz_combit(unsafe_mpz rop, mp_bitcnt_t bit_index);
void pmpz_setbit(const unsafe_mpz op, mp_bitcnt_t bit_index);

// Input and Output Functions
size_t pmpz_inp_raw(unsafe_mpz rop, FILE *stream);
size_t pmpz_inp_str(unsafe_mpz rop, FILE *stream, int base);
size_t pmpz_out_raw(FILE *stream, const unsafe_mpz op);
size_t pmpz_out_str(FILE *stream, int base, const unsafe_mpz op);

// Random Number Functions
unsafe_mpz pmpz_random(mp_size_t max_size);
unsafe_mpz pmpz_random2(mp_size_t max_size);
unsafe_mpz pmpz_rrandomb(gmp_randstate_t state, mp_bitcnt_t n);
unsafe_mpz pmpz_urandomb(gmp_randstate_t state, mp_bitcnt_t n);
unsafe_mpz pmpz_urandomm(gmp_randstate_t state, const unsafe_mpz n);

// Integer Import and Export
unsafe_mpz pmpz_import(size_t count, int order, size_t size, int endian, size_t nails, const void *op);
void *pmpz_export(int order, size_t size, int endian, size_t nails, const unsafe_mpz op);

// Miscellaneous Functions
int pmpz_even_p(const unsafe_mpz op);
int pmpz_fits_sint_p(const unsafe_mpz op);
int pmpz_fits_slong_p(const unsafe_mpz op);
int pmpz_fits_sshort_p(const unsafe_mpz op);
int pmpz_fits_uint_p(const unsafe_mpz op);
int pmpz_fits_ulong_p(const unsafe_mpz op);
int pmpz_fits_ushort_p(const unsafe_mpz op);
size_t pmpz_sizeinbase(const unsafe_mpz op, int base);

// Special Functions
size_t pmpz_size(const unsafe_mpz op);

#endif /* GO_MPZ_H */
