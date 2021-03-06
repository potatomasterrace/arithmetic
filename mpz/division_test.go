package mpz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDivision(t *testing.T) {
	mpz19 := MpzFromString("19", 10)
	mpz16 := MpzFromString("16", 10)
	mpz7 := MpzFromString("7", 10)
	mpz6 := MpzFromString("6", 10)
	mpz4 := MpzFromString("4", 10)
	mpz3 := MpzFromString("3", 10)

	t.Run("MpzCongruent2ExpP", func(t *testing.T) {
		assert.Equal(t, mpz19.Congruent2ExpP(mpz7, 4), false)
		assert.Equal(t, mpz19.Congruent2ExpP(mpz3, 3), true)
	})

	t.Run("MpzCongruentP", func(t *testing.T) {
		assert.Equal(t, mpz19.CongruentP(mpz3, mpz7), false)
		assert.Equal(t, mpz19.CongruentP(mpz7, mpz3), true)
	})

	t.Run("CongruentUiP", func(t *testing.T) {
		assert.Equal(t, mpz19.CongruentUiP(3, 7), false)
		assert.Equal(t, mpz19.CongruentUiP(7, 3), true)
	})

	t.Run("Divisible2ExpP", func(t *testing.T) {
		assert.Equal(t, mpz19.Divisible2ExpP(2), false)
		assert.Equal(t, mpz16.Divisible2ExpP(2), true)
	})

	t.Run("DivisibleP", func(t *testing.T) {
		assert.Equal(t, mpz19.DivisibleP(mpz7), false)
		assert.Equal(t, mpz16.DivisibleP(mpz4), true)
	})

	t.Run("DivisibleUiP", func(t *testing.T) {
		assert.Equal(t, mpz19.DivisibleUiP(7), false)
		assert.Equal(t, mpz16.DivisibleUiP(4), true)
	})

	t.Run("CDivQUi", func(t *testing.T) {
		q, ui := mpz19.CDivQUi(6)
		assert.Equal(t, q.String(), "4")
		assert.Equal(t, ui, uint(5))
		q, ui = mpz16.CDivQUi(3)
		assert.Equal(t, q.String(), "6")
		assert.Equal(t, ui, uint(2))
	})

	t.Run("CDivQRUi", func(t *testing.T) {
		q, r, rUi := mpz19.CDivQRUi(6)
		assert.Equal(t, q.String(), "4")
		assert.Equal(t, r.String(), "-5")
		assert.Equal(t, rUi, uint(5))
		q, r, rUi = mpz16.CDivQRUi(7)
		assert.Equal(t, q.String(), "3")
		assert.Equal(t, r.String(), "-5")
		assert.Equal(t, rUi, uint(5))
	})

	t.Run("CDivRUi", func(t *testing.T) {
		r, rUi := mpz19.CDivRUi(6)
		assert.Equal(t, r.String(), "-5")
		assert.Equal(t, rUi, uint(5))
		r, rUi = mpz16.CDivRUi(7)
		assert.Equal(t, r.String(), "-5")
		assert.Equal(t, rUi, uint(5))
	})
	t.Run("CDivUi", func(t *testing.T) {
		r, rUi := mpz19.CDivUi(6)
		assert.Equal(t, r.String(), "-5")
		assert.Equal(t, rUi, uint(5))
		r, rUi = mpz16.CDivUi(7)
		assert.Equal(t, r.String(), "-5")
		assert.Equal(t, rUi, uint(5))
	})

	t.Run("FDivQUi", func(t *testing.T) {
		q, qUi := mpz19.FDivQUi(6)
		assert.Equal(t, q.String(), "3")
		assert.Equal(t, qUi, uint(1))
		q, qUi = mpz16.FDivQUi(7)
		assert.Equal(t, q.String(), "2")
		assert.Equal(t, qUi, uint(2))
	})

	t.Run("FDivQRUi", func(t *testing.T) {
		q, r, rUi := mpz19.FDivQRUi(6)
		assert.Equal(t, q.String(), "3")
		assert.Equal(t, r.String(), "1")
		assert.Equal(t, rUi, uint(1))
		q, r, rUi = mpz16.FDivQRUi(7)
		assert.Equal(t, q.String(), "2")
		assert.Equal(t, r.String(), "2")
		assert.Equal(t, rUi, uint(2))
	})

	t.Run("FDivRUi", func(t *testing.T) {
		r, rUi := mpz19.FDivRUi(6)
		assert.Equal(t, r.String(), "3")
		assert.Equal(t, rUi, uint(1))
		r, rUi = mpz16.FDivRUi(7)
		assert.Equal(t, r.String(), "2")
		assert.Equal(t, rUi, uint(2))
	})

	t.Run("FDivUi", func(t *testing.T) {
		rUi := mpz19.FDivUi(6)
		assert.Equal(t, rUi, uint(1))
		rUi = mpz16.FDivUi(7)
		assert.Equal(t, rUi, uint(2))
	})

	t.Run("ModUi", func(t *testing.T) {
		r, rUi := mpz19.ModUi(6)
		assert.Equal(t, r.String(), "1")
		assert.Equal(t, rUi, uint(1))
		r, rUi = mpz16.ModUi(7)
		assert.Equal(t, r.String(), "2")
		assert.Equal(t, rUi, uint(2))
	})

	t.Run("TDivQUi", func(t *testing.T) {
		q, qUi := mpz19.TDivQUi(6)
		assert.Equal(t, q.String(), "3")
		assert.Equal(t, qUi, uint(1))
		q, qUi = mpz16.TDivQUi(7)
		assert.Equal(t, q.String(), "2")
		assert.Equal(t, qUi, uint(2))
	})

	t.Run("TDivQRUi", func(t *testing.T) {
		q, r, rUi := mpz19.TDivQRUi(6)
		assert.Equal(t, q.String(), "3")
		assert.Equal(t, r.String(), "1")
		assert.Equal(t, rUi, uint(1))
		q, r, rUi = mpz16.TDivQRUi(7)
		assert.Equal(t, q.String(), "2")
		assert.Equal(t, r.String(), "2")
		assert.Equal(t, rUi, uint(2))
	})

	t.Run("TDivRUi", func(t *testing.T) {
		r, rUi := mpz19.TDivRUi(6)
		assert.Equal(t, r.String(), "3")
		assert.Equal(t, rUi, uint(1))
		r, rUi = mpz16.TDivRUi(7)
		assert.Equal(t, r.String(), "2")
		assert.Equal(t, rUi, uint(2))
	})

	t.Run("TDivUi", func(t *testing.T) {
		rUi := mpz19.TDivUi(6)
		assert.Equal(t, rUi, uint(1))
		rUi = mpz16.TDivUi(7)
		assert.Equal(t, rUi, uint(2))
	})

	t.Run("CDivQR", func(t *testing.T) {
		q, r := mpz19.CDivQR(mpz6)
		assert.Equal(t, q.String(), "4")
		assert.Equal(t, r.String(), "-5")
		q, r = mpz16.CDivQR(mpz6)
		assert.Equal(t, q.String(), "3")
		assert.Equal(t, r.String(), "-2")
	})

	t.Run("FDivQR", func(t *testing.T) {
		q, r := mpz19.FDivQR(mpz6)
		assert.Equal(t, q.String(), "3")
		assert.Equal(t, r.String(), "1")
		q, r = mpz16.FDivQR(mpz6)
		assert.Equal(t, q.String(), "2")
		assert.Equal(t, r.String(), "4")
	})

	t.Run("TDivQR", func(t *testing.T) {
		q, r := mpz19.TDivQR(mpz6)
		assert.Equal(t, q.String(), "3")
		assert.Equal(t, r.String(), "1")
		q, r = mpz16.TDivQR(mpz6)
		assert.Equal(t, q.String(), "2")
		assert.Equal(t, r.String(), "4")
	})

	t.Run("CDivQ", func(t *testing.T) {
		q := mpz19.CDivQ(mpz6)
		assert.Equal(t, q.String(), "4")
		q = mpz16.CDivQ(mpz6)
		assert.Equal(t, q.String(), "3")
	})

	t.Run("CDivQ2Exp", func(t *testing.T) {
		q := mpz19.CDivQ2Exp(2)
		assert.Equal(t, q.String(), "5")
		q = mpz16.CDivQ2Exp(2)
		assert.Equal(t, q.String(), "4")
	})

	t.Run("CDivR", func(t *testing.T) {
		r := mpz19.CDivR(mpz6)
		assert.Equal(t, r.String(), "-5")
		r = mpz16.CDivR(mpz7)
		assert.Equal(t, r.String(), "-5")
	})

	t.Run("CDivR2Exp", func(t *testing.T) {
		q := mpz19.CDivR2Exp(2)
		assert.Equal(t, q.String(), "-1")
		q = mpz16.CDivR2Exp(2)
		assert.Equal(t, q.String(), "0")
	})

	t.Run("CDivExact", func(t *testing.T) {
		q := mpz19.DivExact(mpz6)
		assert.Equal(t, q.String(), "3")
		q = mpz16.DivExact(mpz4)
		assert.Equal(t, q.String(), "4")
	})

	t.Run("DivExactUi", func(t *testing.T) {
		q := mpz19.DivExactUi(6)
		assert.Equal(t, q.String(), "3")
		q = mpz16.DivExactUi(4)
		assert.Equal(t, q.String(), "4")
	})

	t.Run("FDivQ", func(t *testing.T) {
		q := mpz19.FDivQ(mpz6)
		assert.Equal(t, q.String(), "3")
		q = mpz16.FDivQ(mpz7)
		assert.Equal(t, q.String(), "2")
	})

	t.Run("FDivQ2Exp", func(t *testing.T) {
		q := mpz19.FDivQ2Exp(2)
		assert.Equal(t, q.String(), "4")
		q = mpz16.FDivQ2Exp(2)
		assert.Equal(t, q.String(), "4")
	})

	t.Run("FDivR", func(t *testing.T) {
		q := mpz19.FDivR(mpz6)
		assert.Equal(t, q.String(), "3")
		q = mpz16.FDivR(mpz7)
		assert.Equal(t, q.String(), "2")
	})

	t.Run("FDivR2Exp", func(t *testing.T) {
		q := mpz19.FDivR2Exp(6)
		assert.Equal(t, q.String(), "19")
		q = mpz16.FDivR2Exp(7)
		assert.Equal(t, q.String(), "16")
	})

	t.Run("Mod", func(t *testing.T) {
		q := mpz19.Mod(6)
		assert.Equal(t, q.String(), "19")
		q = mpz16.Mod(7)
		assert.Equal(t, q.String(), "16")
	})

	t.Run("TDivQ", func(t *testing.T) {
		q := mpz19.TDivQ(mpz6)
		assert.Equal(t, q.String(), "3")
		q = mpz16.TDivQ(mpz7)
		assert.Equal(t, q.String(), "2")
	})

	t.Run("TDivQ2Exp", func(t *testing.T) {
		q := mpz19.TDivQ2Exp(2)
		assert.Equal(t, q.String(), "4")
		q = mpz16.TDivQ2Exp(2)
		assert.Equal(t, q.String(), "4")
	})

	t.Run("TDivR", func(t *testing.T) {
		q := mpz19.TDivR(mpz6)
		assert.Equal(t, q.String(), "3")
		q = mpz16.TDivR(mpz7)
		assert.Equal(t, q.String(), "2")
	})

	t.Run("TDivR2Exp", func(t *testing.T) {
		q := mpz19.TDivR2Exp(2)
		assert.Equal(t, q.String(), "3")
		q = mpz16.TDivR2Exp(2)
		assert.Equal(t, q.String(), "0")
	})

}
