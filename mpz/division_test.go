package mpz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDivision(t *testing.T) {
	mpz19, err := MpzFromString("19", 10)
	assert.Nil(t, err)
	mpz16, err := MpzFromString("16", 10)
	assert.Nil(t, err)
	mpz7, err := MpzFromString("7", 10)
	assert.Nil(t, err)
	mpz4, err := MpzFromString("4", 10)
	assert.Nil(t, err)
	mpz3, err := MpzFromString("3", 10)
	assert.Nil(t, err)

	t.Run(("MpzCongruent2ExpP"), func(t *testing.T) {
		assert.Equal(t, mpz19.Congruent2ExpP(mpz7, 4), false)
		assert.Equal(t, mpz19.Congruent2ExpP(mpz3, 3), true)
	})

	t.Run(("MpzCongruentP"), func(t *testing.T) {
		assert.Equal(t, mpz19.CongruentP(mpz3, mpz7), false)
		assert.Equal(t, mpz19.CongruentP(mpz7, mpz3), true)
	})

	t.Run(("CongruentUiP"), func(t *testing.T) {
		assert.Equal(t, mpz19.CongruentUiP(3, 7), false)
		assert.Equal(t, mpz19.CongruentUiP(7, 3), true)
	})

	t.Run(("Divisible2ExpP"), func(t *testing.T) {
		assert.Equal(t, mpz19.Divisible2ExpP(2), false)
		assert.Equal(t, mpz16.Divisible2ExpP(2), true)
	})

	t.Run(("DivisibleP"), func(t *testing.T) {
		assert.Equal(t, mpz19.DivisibleP(mpz7), false)
		assert.Equal(t, mpz16.DivisibleP(mpz4), true)
	})

	t.Run(("DivisibleUiP"), func(t *testing.T) {
		assert.Equal(t, mpz19.DivisibleUiP(7), false)
		assert.Equal(t, mpz16.DivisibleUiP(4), true)
	})

	t.Run(("CDivQUi"), func(t *testing.T) {
		q, ui := mpz19.CDivQUi(6)
		assert.Equal(t, q.StringOrPanic(), "4")
		assert.Equal(t, ui, uint(5))
		q, ui = mpz16.CDivQUi(3)
		assert.Equal(t, q.StringOrPanic(), "6")
		assert.Equal(t, ui, uint(2))
	})

	t.Run(("CDivQRUi"), func(t *testing.T) {
		q, r, rUi := mpz19.CDivQRUi(6)
		assert.Equal(t, q.StringOrPanic(), "4")
		assert.Equal(t, r.StringOrPanic(), "-5")
		assert.Equal(t, rUi, uint(5))
		q, r, rUi = mpz16.CDivQRUi(7)
		assert.Equal(t, q.StringOrPanic(), "3")
		assert.Equal(t, r.StringOrPanic(), "-5")
		assert.Equal(t, rUi, uint(5))
	})

	t.Run(("CDivRUi"), func(t *testing.T) {
		q, r := mpz19.CDivRUi(6)
		assert.Equal(t, q.StringOrPanic(), "-5")
		assert.Equal(t, r, uint(5))
		q, r = mpz16.CDivRUi(9)
		assert.Equal(t, q.StringOrPanic(), "-2")
		assert.Equal(t, r, uint(2))
		q, r = mpz16.CDivRUi(8)
		assert.Equal(t, q.StringOrPanic(), "0")
		assert.Equal(t, r, uint(0))
	})

	t.Run(("CDivUi"), func(t *testing.T) {
		r := mpz19.CDivUi(6)
		assert.Equal(t, r, uint(5))
		r = mpz16.CDivUi(9)
		assert.Equal(t, r, uint(2))
		r = mpz16.CDivUi(8)
		assert.Equal(t, r, uint(0))
	})

	t.Run(("FDivQUi"), func(t *testing.T) {
		q, r := mpz19.FDivQUi(6)
		assert.Equal(t, q.StringOrPanic(), "3")
		assert.Equal(t, r, uint(1))
		q, r = mpz16.FDivQUi(7)
		assert.Equal(t, q.StringOrPanic(), "2")
		assert.Equal(t, r, uint(2))
	})

}
