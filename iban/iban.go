package iban

import (
	"errors"
	"math/big"
	"strconv"
	"strings"

	"golang.org/x/exp/utf8string" // TODO: resolve this dependency by implementing this in codebase
)

var InvalidLength = errors.New("Invalid Iban account length")
var InvalidRemainder = errors.New("Invalid Iban remainder")
var LogicalError = errors.New("Unexpected error, can't process")
var InvalidEncoding = errors.New("Invalid character encoding")
var InvalidCharacter = errors.New("Invalid character in account string")

var bigNintySeven = big.NewInt(97)
var bigOne = big.NewInt(1)

func asciiAlphaNumToNum(asciiNumbers []byte) ([]byte, error) {
	out := []byte{}
	for _, c := range asciiNumbers {
		if c >= '0' && c <= '9' {
			// Already in ASCII
			out = append(out, c)
		} else if c >= 'A' && c <= 'Z' {
			// Convert the letter into a numerical value, then generate the
			// string that represents the value in base 10. Finally append the
			// ASCII code points.
			numValue := int(c-'A') + 10
			stringValue := strconv.Itoa(numValue)
			out = append(out, []byte(stringValue)...)
		} else {
			return nil, InvalidCharacter
		}
	}
	return out, nil
}

type Iban struct {
	raw           string
	accountNumber *big.Int
}

func (i *Iban) ValidationError() error {
	if !i.validLength() {
		return InvalidLength
	}

	okRemainder, err := i.validMod97()
	if err != nil {
		return err
	}

	if !okRemainder {
		return InvalidRemainder
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

func (i *Iban) validMod97() (bool, error) {
	var remainder = &big.Int{}
	asNum, err := i.asNumber()

	if err != nil {
		return false, err
	}

	remainder.Mod(asNum, bigNintySeven)
	return (remainder.Cmp(bigOne) == 0), nil
}

func (i *Iban) asNumber() (*big.Int, error) {
	if i.accountNumber != nil {
		return i.accountNumber, nil
	}

	if !utf8string.NewString(i.raw).IsASCII() {
		return nil, InvalidEncoding
	}

	rawBytes := []byte(i.raw)
	numberBuffer := []byte{}

	accountNumber, err := asciiAlphaNumToNum(rawBytes[4:])
	if err != nil {
		return nil, err
	}
	numberBuffer = append(numberBuffer, accountNumber...)

	countryCode, err := asciiAlphaNumToNum(rawBytes[0:2])

	if err != nil {
		return nil, err
	}
	numberBuffer = append(numberBuffer, countryCode...)

	remainderCorrection, err := asciiAlphaNumToNum(rawBytes[2:4])
	if err != nil {
		return nil, err
	}
	numberBuffer = append(numberBuffer, remainderCorrection...)

	n := new(big.Int)
	n, ok := n.SetString(string(numberBuffer), 10)
	if !ok {
		return nil, LogicalError
	}
	return n, nil
}

func (i *Iban) String() string {
	return i.raw
}

func New(content string) *Iban {
	return &Iban{
		raw: strings.ToUpper(content),
	}
}
