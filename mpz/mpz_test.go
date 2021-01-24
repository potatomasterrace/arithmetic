package mpz

import (
	"fmt"
	"runtime"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

// Test that ptr returns the pointer value.
func TestPtr(t *testing.T) {
	testvar := ""
	ptr := unsafe.Pointer(&testvar)
	mpz := Mpz{
		ptr: &ptr,
	}
	assert.Equal(t, mpz.Ptr(), ptr)
}

// Test that garbage collection is done.
func TestDeferencing(t *testing.T) {
	var allocatedPtr unsafe.Pointer
	var deallocatedPtr unsafe.Pointer
	oldClearMpz := clearMpz
	stubDeallocation := func(ptr *unsafe.Pointer) {
		fmt.Println("deallocated")
		deallocatedPtr = *ptr
	}
	clearMpz = stubDeallocation
	someAllocation := func() {
		testvar := ""
		ptr := unsafe.Pointer(&testvar)
		allocatedPtr = ptr
		mpzFromPtr(ptr)
	}
	someAllocation()
	for i := 0; i < 10; i++ {
		runtime.GC()
	}
	clearMpz = oldClearMpz
	assert.Equal(t, allocatedPtr, deallocatedPtr)
}
