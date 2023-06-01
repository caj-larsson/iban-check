package iban_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/caj-larsson/iban-check/v2/iban"
	"github.com/matryer/is"
)

func TestLength(t *testing.T) {
	is := is.New(t)
	// TODO: refactor into table
	is.True(iban.New("SE950").ValidationError() != iban.InvalidLength)
	is.True(errors.Is(iban.New("SE01").ValidationError(), iban.InvalidLength))
	is.True(errors.Is(iban.New("SE01"+strings.Repeat("0", 31)).ValidationError(), iban.InvalidLength))
}

func TestModulo(t *testing.T) {
	is := is.New(t)

	is.Equal(iban.New("SE950").ValidationError(), nil)
	// wikipedia example
	is.Equal(iban.New("GB82WEST12345698765432").ValidationError(), nil)
	is.True(errors.Is(iban.New("SE010").ValidationError(), iban.InvalidRemainder))
}

func TestCasing(t *testing.T) {
	is := is.New(t)
	// TODO: refactor into table
	is.Equal(iban.New("SE950").ValidationError(), nil)
	is.Equal(iban.New("se950").ValidationError(), nil)
}
