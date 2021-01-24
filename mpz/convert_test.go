package mpz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertion(t *testing.T) {
	mpz100, err := MpzFromString("100", 10)
	assert.Nil(t, err)
	mpz12, err := MpzFromString("-12", 10)
	assert.Nil(t, err)

	t.Run(("GetString"), func(t *testing.T) {
		assert.Nil(t, err)
		assert.Equal(t, mpz100.GetStringOrPanic(16), "64")
		assert.Equal(t, mpz12.GetStringOrPanic(16), "-c")
	})

	t.Run(("String"), func(t *testing.T) {
		assert.Nil(t, err)
		assert.Equal(t, mpz100.StringOrPanic(), "100")
		assert.Equal(t, mpz12.StringOrPanic(), "-12")
	})
	t.Run(("GetDouble"), func(t *testing.T) {
		assert.Nil(t, err)
		assert.Equal(t, mpz100.GetDoubleOrPanic(), float64(100))
		assert.Equal(t, mpz12.GetDoubleOrPanic(), float64(-12))
	})

	t.Run(("GetD2Exp"), func(t *testing.T) {
		v_100, e_100 := mpz100.GetD2ExpOrPanic()
		assert.Equal(t, v_100, float64(0.78125))
		assert.Equal(t, e_100, int(7))

		v_12, e_12 := mpz12.GetD2ExpOrPanic()
		assert.Equal(t, v_12, float64(-0.75))
		assert.Equal(t, e_12, int(4))
	})

	t.Run("GetSignedInt", func(t *testing.T) {
		i100, i12 := mpz100.GetSignedIntOrPanic(), mpz12.GetSignedIntOrPanic()
		assert.Equal(t, i100, int(100))
		assert.Equal(t, i12, int(-12))
	})

	t.Run("GetUnSignedInt", func(t *testing.T) {
		i100, i12 := mpz100.GetUnsignedIntOrPanic(), mpz12.GetUnsignedIntOrPanic()
		assert.Equal(t, i100, uint(100))
		assert.Equal(t, i12, uint(12))
	})

}
