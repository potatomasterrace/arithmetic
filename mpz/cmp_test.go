package mpz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Arithmetic Functions
func TestCmp(t *testing.T) {
	mpz100, err := MpzFromString("100", 10)
	assert.Nil(t, err)
	mpz8, err := MpzFromString("8", 10)
	assert.Nil(t, err)
	minusMpz8, err := MpzFromString("-8", 10)
	assert.Nil(t, err)
	t.Run(("cmp"), func(t *testing.T) {
		assert.Equal(t, mpz8.CmpOrPanic(mpz100), -1)
		assert.Equal(t, mpz8.CmpOrPanic(mpz8), 0)
		assert.Equal(t, mpz100.CmpOrPanic(mpz8), 1)
	})
	t.Run(("cmpD"), func(t *testing.T) {
		assert.Equal(t, mpz8.CmpDOrPanic(100), -1)
		assert.Equal(t, mpz8.CmpDOrPanic(-12), 1)
		assert.Equal(t, mpz100.CmpDOrPanic(-12), 1)
	})
	t.Run(("cmpSi"), func(t *testing.T) {
		assert.Equal(t, mpz100.CmpSiOrPanic(101), -1)
		assert.Equal(t, mpz100.CmpSiOrPanic(100), 0)
		assert.Equal(t, mpz100.CmpSiOrPanic(-99), 2)
	})
	t.Run(("cmpUi"), func(t *testing.T) {
		assert.Equal(t, mpz100.CmpUiOrPanic(101), -1)
		assert.Equal(t, mpz100.CmpUiOrPanic(100), 0)
		assert.Equal(t, mpz100.CmpUiOrPanic(99), 1)
	})
	t.Run(("cmpAbs"), func(t *testing.T) {
		assert.Equal(t, minusMpz8.CmpAbsOrPanic(mpz100), -1)
		assert.Equal(t, minusMpz8.CmpAbsOrPanic(mpz8), 0)
		assert.Equal(t, mpz100.CmpAbsOrPanic(minusMpz8), 1)

		assert.Equal(t, mpz8.CmpAbsOrPanic(mpz100), -1)
		assert.Equal(t, mpz8.CmpAbsOrPanic(minusMpz8), 0)
		assert.Equal(t, mpz100.CmpAbsOrPanic(minusMpz8), 1)
	})
	t.Run(("cmpAbsD"), func(t *testing.T) {
		assert.Equal(t, minusMpz8.CmpAbsDOrPanic(8), -1)
		assert.Equal(t, mpz100.CmpAbsDOrPanic(100), 0)
		assert.Equal(t, mpz100.CmpAbsDOrPanic(-12), -1)

		assert.Equal(t, mpz8.CmpAbsDOrPanic(12), -1)
		assert.Equal(t, mpz8.CmpAbsDOrPanic(8), 0)
		assert.Equal(t, mpz100.CmpAbsDOrPanic(-12), -1)
	})

	t.Run(("cmpAbsUi"), func(t *testing.T) {
		assert.Equal(t, mpz100.CmpAbsUiOrPanic(99), 1)
		assert.Equal(t, mpz100.CmpAbsUiOrPanic(100), 0)
		assert.Equal(t, mpz100.CmpAbsUiOrPanic(101), -1)

		assert.Equal(t, minusMpz8.CmpAbsUiOrPanic(7), -1)
		assert.Equal(t, minusMpz8.CmpAbsUiOrPanic(9), -1)
	})

	t.Run(("sign"), func(t *testing.T) {
		minusMpz0, err := MpzFromString("0", 10)
		assert.Nil(t, err)
		assert.Equal(t, minusMpz8.SignOrPanic(), -1)
		assert.Equal(t, mpz100.SignOrPanic(), 1)
		assert.Equal(t, minusMpz0.SignOrPanic(), 0)
	})
}
