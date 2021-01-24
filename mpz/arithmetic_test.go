package mpz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Arithmetic Functions

func TestAdd(t *testing.T) {

	mpz100 := MpzFromString("100", 10)
	mpz12 := MpzFromString("-12", 10)
	t.Run(("add"), func(t *testing.T) {
		mpz := mpz12.Add(mpz100)

		assert.Equal(t, mpz.String(), "88")
	})
	t.Run(("addMany"), func(t *testing.T) {
		mpz := mpz12.AddMany(mpz100, mpz12)

		assert.Equal(t, mpz.String(), "76")
	})

	t.Run(("addUi"), func(t *testing.T) {
		mpz := mpz12.AddUi(51)

		assert.Equal(t, mpz.String(), "39")
	})
}

func TestSub(t *testing.T) {

	mpz100 := MpzFromString("100", 10)
	mpz12 := MpzFromString("-12", 10)
	t.Run(("sub"), func(t *testing.T) {
		mpz := mpz12.Sub(mpz100)

		assert.Equal(t, mpz.String(), "-112")
	})
	t.Run(("subMany"), func(t *testing.T) {
		mpz := mpz12.SubMany(mpz100, mpz100)

		assert.Equal(t, mpz.String(), "-212")
	})

	t.Run(("subUi"), func(t *testing.T) {
		mpz := mpz100.SubUi(8)

		assert.Equal(t, mpz.String(), "92")
	})

	t.Run(("UiSub"), func(t *testing.T) {
		mpz := mpz100.UiSub(8)

		assert.Equal(t, mpz.String(), "-92")
	})

}

func TestMul(t *testing.T) {
	mpz100 := MpzFromString("100", 10)
	mpz12 := MpzFromString("-12", 10)
	t.Run(("mul"), func(t *testing.T) {
		mpz := mpz12.Mul(mpz100)

		assert.Equal(t, mpz.String(), "-1200")
	})
	t.Run(("mulMany"), func(t *testing.T) {
		mpz := mpz12.MulMany(mpz12, mpz100)

		assert.Equal(t, mpz.String(), "14400")
	})
	t.Run(("mulSi"), func(t *testing.T) {
		mpz := mpz12.MulSi(-10)

		assert.Equal(t, mpz.String(), "120")
	})
	t.Run(("mulUi"), func(t *testing.T) {
		mpz := mpz12.MulUi(10)

		assert.Equal(t, mpz.String(), "-120")
	})

}

func TestAddMul(t *testing.T) {
	mpz100 := MpzFromString("100", 10)
	mpz12 := MpzFromString("-12", 10)
	mpz7 := MpzFromString("7", 10)

	t.Run(("AddMul"), func(t *testing.T) {
		mpz := mpz12.AddMul(mpz100, mpz7)

		assert.Equal(t, mpz.String(), "688")
	})

	t.Run(("AddMulUi"), func(t *testing.T) {
		mpz := mpz12.AddMulUi(mpz100, 7)

		assert.Equal(t, mpz.String(), "688")
	})
}

func TestSubMul(t *testing.T) {
	mpz100 := MpzFromString("100", 10)
	mpz12 := MpzFromString("-12", 10)
	mpz7 := MpzFromString("7", 10)

	t.Run(("subMul"), func(t *testing.T) {
		mpz := mpz12.SubMul(mpz100, mpz7)

		assert.Equal(t, mpz.String(), "-712")
	})

	t.Run(("subMulUi"), func(t *testing.T) {
		mpz := mpz12.SubMulUi(mpz100, 7)

		assert.Equal(t, mpz.String(), "-712")
	})
}

func TestMul2Exp(t *testing.T) {
	mpz := MpzFromString("-100", 10)
	assert.Equal(t, mpz.String(), "-100")
	mpzAbs := mpz.Mul2Exp(2)
	assert.Equal(t, mpzAbs.String(), "-400")
}
func TestAbs(t *testing.T) {
	mpz := MpzFromString("-100", 10)
	assert.Equal(t, mpz.String(), "-100")
	mpzAbs := mpz.Abs()
	assert.Equal(t, mpzAbs.String(), "100")
}

func TestNeg(t *testing.T) {
	mpz := MpzFromString("-100", 10)
	assert.Equal(t, mpz.String(), "-100")
	mpzAbs := mpz.Neg()
	assert.Equal(t, mpzAbs.String(), "100")
}
