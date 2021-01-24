package mpz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Arithmetic Functions
func TestCmp(t *testing.T) {
	mpz100 := MpzFromString("100", 10)

	mpz8 := MpzFromString("8", 10)

	minusMpz8 := MpzFromString("-8", 10)

	t.Run(("cmp"), func(t *testing.T) {
		assert.Equal(t, mpz8.Cmp(mpz100), -1)
		assert.Equal(t, mpz8.Cmp(mpz8), 0)
		assert.Equal(t, mpz100.Cmp(mpz8), 1)
	})
	t.Run(("cmpD"), func(t *testing.T) {
		assert.Equal(t, mpz8.CmpD(100), -1)
		assert.Equal(t, mpz8.CmpD(-12), 1)
		assert.Equal(t, mpz100.CmpD(-12), 1)
	})
	t.Run(("cmpSi"), func(t *testing.T) {
		assert.Equal(t, mpz100.CmpSi(101), -1)
		assert.Equal(t, mpz100.CmpSi(100), 0)
		assert.Equal(t, mpz100.CmpSi(-99), 2)
	})
	t.Run(("cmpUi"), func(t *testing.T) {
		assert.Equal(t, mpz100.CmpUi(101), -1)
		assert.Equal(t, mpz100.CmpUi(100), 0)
		assert.Equal(t, mpz100.CmpUi(99), 1)
	})
	t.Run(("cmpAbs"), func(t *testing.T) {
		assert.Equal(t, minusMpz8.CmpAbs(mpz100), -1)
		assert.Equal(t, minusMpz8.CmpAbs(mpz8), 0)
		assert.Equal(t, mpz100.CmpAbs(minusMpz8), 1)

		assert.Equal(t, mpz8.CmpAbs(mpz100), -1)
		assert.Equal(t, mpz8.CmpAbs(minusMpz8), 0)
		assert.Equal(t, mpz100.CmpAbs(minusMpz8), 1)
	})
	t.Run(("cmpAbsD"), func(t *testing.T) {
		assert.Equal(t, minusMpz8.CmpAbsD(8), -1)
		assert.Equal(t, mpz100.CmpAbsD(100), 0)
		assert.Equal(t, mpz100.CmpAbsD(-12), -1)

		assert.Equal(t, mpz8.CmpAbsD(12), -1)
		assert.Equal(t, mpz8.CmpAbsD(8), 0)
		assert.Equal(t, mpz100.CmpAbsD(-12), -1)
	})

	t.Run(("cmpAbsUi"), func(t *testing.T) {
		assert.Equal(t, mpz100.CmpAbsUi(99), 1)
		assert.Equal(t, mpz100.CmpAbsUi(100), 0)
		assert.Equal(t, mpz100.CmpAbsUi(101), -1)

		assert.Equal(t, minusMpz8.CmpAbsUi(7), -1)
		assert.Equal(t, minusMpz8.CmpAbsUi(9), -1)
	})

	t.Run(("sign"), func(t *testing.T) {
		minusMpz0 := MpzFromString("0", 10)

		assert.Equal(t, minusMpz8.Sign(), -1)
		assert.Equal(t, mpz100.Sign(), 1)
		assert.Equal(t, minusMpz0.Sign(), 0)
	})
}
