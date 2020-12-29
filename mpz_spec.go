package main

// Special Functions
// size_t pmpz_size(const unsafe_mpz op);

// #cgo LDFLAGS: -lgmp
// #include "mpz.h"
import "C"

func (m Mpz) Size() C.ulong {
	return C.pmpz_size(m.Ptr())
}
