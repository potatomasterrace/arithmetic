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
    MALLOC_PROXY(p_mpz,"mpz",value,sizeof(mpz_t)); \
    /* initialize mpz pointer  */        \
    mpz_init((mpz_t * )value);

// Macro for wrapping mpz functions
#define wrap_mpz_function(func, ...) \
    /* init return reference*/       \
    p_mpz_init(rop);                 \
    /* Call func */                  \
    func(*rop, __VA_ARGS__);         \
    /* return casted value */        \
    return (unsafe_mpz)rop;
