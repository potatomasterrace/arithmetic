#ifndef GO_RANDSTATE_H
#define GO_RANDSTATE_H

#include <gmp.h>
#include <malloc.h>

#define p_randstate gmp_randstate_t *

#define ref_to_gmp(ref) \
    p_randstate p_##ref = (p_randstate)ref;

#define ref(value) \
    *p_##value

#define unsafe_randstate void *

// Macro for initializing mpz
#define p_randstate_init(value)                \
    /* allocate memory for reference */  \
    MALLOC_PROXY(p_randstate,"rnd",value,sizeof(gmp_randstate_t)); \
    /* initialize mpz pointer  */        \
    mpz_init(value);


#endif /* GO_RANDSTATE_H */