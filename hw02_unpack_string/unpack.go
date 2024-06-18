package main

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	if str == "" {
		return "", nil
	}

	builder := strings.Builder{}
	var ch string
	prevWasDigit := false

	// check if the 1st character is digit
	for i, r := range str {
		if i == 0 && unicode.IsDigit(r) {
			return "", ErrInvalidString
		}

		// if its number
		if unicode.IsDigit(r) {
			if prevWasDigit {
				return "", ErrInvalidString
			}

			n, err := strconv.Atoi(string(r))
			if err != nil {
				return "", ErrInvalidString
			}

			builder.WriteString(strings.Repeat(ch, n))
			ch = ""
			prevWasDigit = true
		} else { // if its character
			if ch != "" {
				builder.WriteString(ch)
			}
			ch = string(r)
			prevWasDigit = false
		}
	}

	// Add the last character if it wasn't followed by a number
	if ch != "" {
		builder.WriteString(ch)
	}

	return builder.String(), nil
}
