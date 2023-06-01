package iban

import (
	"errors"
	"math/big"
)



type Iban struct {
	raw string
	accountNumber *big.Int
}

var InvalidLength = errors.New("Invalid Iban account length")

var bigNintySeven = big.NewInt(97)
var bigOne = big.NewInt(1)


func (i *Iban) ValidationError() *error {
	if (!i.validLength()) {
		return &InvalidLength
	}

	return nil
}

func (i *Iban) validLength() bool {
	// Minimal possible length is:\
	//  2 letter countrycode
	//  2 letter modulo adjustment
	//  1 letter account number?
	// Max accountnumber length is 30
	//
	nLetters := len(i.raw)
	return (nLetters > 4 && nLetters <= 34)
}

func (i *Iban) validMod97() bool {
	var remainder = &big.Int{}
	remainder.Mod(i.accountNumber, bigNintySeven)
	return remainder.Cmp(bigOne) == 0
}


func (i *Iban) String() string {
	return i.raw
}

func New(asText string) *Iban {
	return &Iban {
		raw: asText,
	}
}
