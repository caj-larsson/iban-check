package iban_test

import (
	"errors"
	"strings"
	"testing"

	"git.sg.caj.me/caj/iban-check/v2/iban"
	"github.com/matryer/is"
)

func TestLength(t *testing.T) {
	is := is.New(t)
	// TODO: refactor into table
	is.Equal(iban.New("SE001").ValidationError(), nil)
	is.True(errors.Is(*iban.New("SE01").ValidationError(), iban.InvalidLength))
	is.True(errors.Is(*iban.New("SE01"+ strings.Repeat("0", 31)).ValidationError(), iban.InvalidLength))
}
