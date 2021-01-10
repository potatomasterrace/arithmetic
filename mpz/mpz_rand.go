package mpz

// #include "mpz.h"
import "C"

// Random Number Functions
// Standby
// unsafe_mpz pmpz_random(mp_size_t max_size);
func MpzRandom(maxSize uint64) Mpz {
	ptr := C.pmpz_random(C.long(maxSize))
	return mpzFromPtr(ptr)
}

// unsafe_mpz pmpz_random2(mp_size_t max_size);
func MpzRandom2(maxSize uint64) Mpz {
	ptr := C.pmpz_random2(C.long(maxSize))
	return mpzFromPtr(ptr)
}

// unsafe_mpz pmpz_rrandomb(gmp_randstate_t state, mp_bitcnt_t n);
// unsafe_mpz pmpz_urandomb(gmp_randstate_t state, mp_bitcnt_t n);
// unsafe_mpz pmpz_urandomm(gmp_randstate_t state, const unsafe_mpz n);
