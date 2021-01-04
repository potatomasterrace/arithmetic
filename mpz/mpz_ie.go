package mpz

// #cgo LDFLAGS: -lgmp
// #include "mpz.h"

// Integer Import and Export
// unsafe_mpz pmpz_import(size_t count, int order, size_t size, int endian, size_t nails, const void *op);

// void *pmpz_export(int order, size_t size, int endian, size_t nails, const unsafe_mpz op);
