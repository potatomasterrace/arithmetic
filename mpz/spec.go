package mpz

// Special Functions
// size_t pmpz_size(const unsafe_mpz op);

// #include "mpz.h"
import "C"

func (op Mpz) Size() C.size_t {
	return C.pmpz_size(op.Ptr())
}

// size_t pmpz_sizeinbase(const unsafe_mpz op, int base);
func (op Mpz) SizeInBase(base C.int) C.size_t {
	return C.pmpz_sizeinbase(op.Ptr(), base)
}
