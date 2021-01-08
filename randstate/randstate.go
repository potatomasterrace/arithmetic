package randstate

// #cgo LDFLAGS: -lgmp
// #include <gmp.h>
import "C"

type RandState *C.gmp_randstate_t

func CreateRandState() RandState {

}
