package mpz

// #cgo LDFLAGS: -lgmp
// #include "mpz.h"
import "C"

// Input and Output Functions
// size_t pmpz_inp_raw(unsafe_mpz rop, FILE *stream);
func MpzInpRaw(stream *C.FILE) (rop Mpz, size C.size_t) {
	rop = MpzInit()
	size = C.pmpz_inp_raw(rop.Ptr(), stream)
	return rop, size
}

// size_t pmpz_inp_str(unsafe_mpz rop, FILE *stream, int base);
func MpzInpStr(stream *C.FILE, base C.int) (rop Mpz, size C.size_t) {
	rop = MpzInit()
	size = C.pmpz_inp_str(rop.Ptr(), stream, base)
	return rop, size
}

// size_t pmpz_out_raw(FILE *stream, const unsafe_mpz op);
func (op Mpz) MpzOutRaw(stream *C.FILE) C.size_t {
	return C.pmpz_out_raw(stream, op.Ptr())
}

// size_t pmpz_out_str(FILE *stream, int base, const unsafe_mpz op);
func (op Mpz) MpzOutStr(stream *C.FILE, base C.int) C.size_t {
	return C.pmpz_out_str(stream, base, op.Ptr())
}
