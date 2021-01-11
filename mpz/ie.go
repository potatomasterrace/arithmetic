package mpz

import "unsafe"

// #include "mpz.h"
import "C"

// Integer Import and Export
// unsafe_mpz pmpz_import(size_t count, int order, size_t size, int endian, size_t nails, const void *op);
func MpzImport(count C.size_t, order C.int, size C.size_t, endian C.int, nails C.size_t, op unsafe.Pointer) (Mpz, error) {
	ptr := C.pmpz_import(count, order, size, endian, nails, op)
	return mpzFromPtr(ptr)
}

// void *pmpz_export(int order, size_t size, int endian, size_t nails, const unsafe_mpz op);
func MpzExport(order C.int, size C.size_t, endian C.int, nails C.size_t, op Mpz) unsafe.Pointer {
	return C.pmpz_export(order, size, endian, nails, op.Ptr())
}
