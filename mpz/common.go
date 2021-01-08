package mpz

import (
	"fmt"
	"runtime"
	"unsafe"
)

// #include <malloc.h>
import "C"

func free(ptr unsafe.Pointer) {
	go C.free(ptr)
}

func RegisterPtrReference(ptr *unsafe.Pointer, callback func(collectedPtr *unsafe.Pointer)) {
	fmt.Println("allocated ptr at : ", *ptr)
	runtime.SetFinalizer(ptr, func(collectedPtr *unsafe.Pointer) {
		fmt.Println("freeing ptr at : ", *collectedPtr)
		go callback(collectedPtr)
	})
}
