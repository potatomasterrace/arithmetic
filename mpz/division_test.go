package mpz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDivision(t *testing.T) {
	mpz19, err := MpzFromString("19", 10)
	assert.Nil(t, err)
	mpz7, err := MpzFromString("7", 10)
	assert.Nil(t, err)
	mpz3, err := MpzFromString("3", 10)
	assert.Nil(t, err)

	t.Run(("MpzCongruent2ExpP"), func(t *testing.T) {
		assert.Equal(t, mpz19.Congruent2ExpP(mpz7, 4), 0)
		assert.Equal(t, mpz19.Congruent2ExpP(mpz3, 3), 1)
	})

	t.Run(("MpzCongruentP"), func(t *testing.T) {
		assert.Equal(t, mpz19.CongruentP(mpz3, mpz7), 0)
		assert.Equal(t, mpz19.CongruentP(mpz7, mpz3), 1)
	})

}
