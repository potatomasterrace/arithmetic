package mpz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Arithmetic Functions

func TestAdd(t *testing.T) {

	mpz100, err := MpzFromString("100", 10)
	assert.Nil(t, err)
	mpz12, err := MpzFromString("-12", 10)
	assert.Nil(t, err)
	t.Run(("add"), func(t *testing.T) {
		mpz, err := mpz12.Add(mpz100)
		assert.Nil(t, err)
		assert.Equal(t, mpz.String(), "88")
	})
	t.Run(("addMany"), func(t *testing.T) {
		mpz, err := mpz12.AddMany(mpz100, mpz12)
		assert.Nil(t, err)
		assert.Equal(t, mpz.String(), "76")
	})

	t.Run(("addUi"), func(t *testing.T) {
		mpz, err := mpz12.AddUi(51)
		assert.Nil(t, err)
		assert.Equal(t, mpz.String(), "39")
	})
}

func TestSub(t *testing.T) {

	mpz100, err := MpzFromString("100", 10)
	assert.Nil(t, err)
	mpz12, err := MpzFromString("-12", 10)
	assert.Nil(t, err)
	t.Run(("sub"), func(t *testing.T) {
		mpz, err := mpz12.Sub(mpz100)
		assert.Nil(t, err)
		assert.Equal(t, mpz.String(), "-112")
	})
	t.Run(("subMany"), func(t *testing.T) {
		mpz, err := mpz12.SubMany(mpz100, mpz100)
		assert.Nil(t, err)
		assert.Equal(t, mpz.String(), "-212")
	})

	t.Run(("subUi"), func(t *testing.T) {
		mpz, err := mpz100.SubUi(8)
		assert.Nil(t, err)
		assert.Equal(t, mpz.String(), "92")
	})

	t.Run(("UiSub"), func(t *testing.T) {
		mpz, err := mpz100.UiSub(8)
		assert.Nil(t, err)
		assert.Equal(t, mpz.String(), "-92")
	})

}

func TestMul(t *testing.T) {
	mpz100, err := MpzFromString("100", 10)
	assert.Nil(t, err)
	mpz12, err := MpzFromString("-12", 10)
	assert.Nil(t, err)
	t.Run(("mul"), func(t *testing.T) {
		mpz, err := mpz12.Mul(mpz100)
		assert.Nil(t, err)
		assert.Equal(t, mpz.String(), "-1200")
	})
	t.Run(("mulMany"), func(t *testing.T) {
		mpz, err := mpz12.MulMany(mpz12, mpz100)
		assert.Nil(t, err)
		assert.Equal(t, mpz.String(), "14400")
	})
	t.Run(("mulSi"), func(t *testing.T) {
		mpz, err := mpz12.MulSi(-10)
		assert.Nil(t, err)
		assert.Equal(t, mpz.String(), "120")
	})
	t.Run(("mulUi"), func(t *testing.T) {
		mpz, err := mpz12.MulUi(10)
		assert.Nil(t, err)
		assert.Equal(t, mpz.String(), "-120")
	})

}

func TestAddMul(t *testing.T) {
	mpz100, err := MpzFromString("100", 10)
	assert.Nil(t, err)
	mpz12, err := MpzFromString("-12", 10)
	assert.Nil(t, err)
	mpz7, err := MpzFromString("7", 10)
	assert.Nil(t, err)

	t.Run(("AddMul"), func(t *testing.T) {
		mpz, err := mpz12.AddMul(mpz100, mpz7)
		assert.Nil(t, err)
		assert.Equal(t, mpz.String(), "688")
	})

	t.Run(("AddMulUi"), func(t *testing.T) {
		mpz, err := mpz12.AddMulUi(mpz100, 7)
		assert.Nil(t, err)
		assert.Equal(t, mpz.String(), "688")
	})
}

func TestSubMul(t *testing.T) {
	mpz100, err := MpzFromString("100", 10)
	assert.Nil(t, err)
	mpz12, err := MpzFromString("-12", 10)
	assert.Nil(t, err)
	mpz7, err := MpzFromString("7", 10)
	assert.Nil(t, err)

	t.Run(("subMul"), func(t *testing.T) {
		mpz, err := mpz12.SubMul(mpz100, mpz7)
		assert.Nil(t, err)
		assert.Equal(t, mpz.String(), "-712")
	})

	t.Run(("subMulUi"), func(t *testing.T) {
		mpz, err := mpz12.SubMulUi(mpz100, 7)
		assert.Nil(t, err)
		assert.Equal(t, mpz.String(), "-712")
	})
}

func TestMul2Exp(t *testing.T) {
	mpz, err := MpzFromString("-100", 10)
	assert.Nil(t, err)
	assert.Equal(t, mpz.String(), "-100")
	mpzAbs, err := mpz.Mul2Exp(2)
	assert.Nil(t, err)
	assert.Equal(t, mpzAbs.String(), "-400")
}
func TestAbs(t *testing.T) {
	mpz, err := MpzFromString("-100", 10)
	assert.Nil(t, err)
	assert.Equal(t, mpz.String(), "-100")
	mpzAbs, err := mpz.Abs()
	assert.Nil(t, err)
	assert.Equal(t, mpzAbs.String(), "100")
}

func TestNeg(t *testing.T) {
	mpz, err := MpzFromString("-100", 10)
	assert.Nil(t, err)
	assert.Equal(t, mpz.String(), "-100")
	mpzAbs, err := mpz.Neg()
	assert.Nil(t, err)
	assert.Equal(t, mpzAbs.String(), "100")
}
