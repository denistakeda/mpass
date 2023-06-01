package client

import (
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/pkg/errors"
)

type (
	paramReader struct {
		printer printer
		scanner scanner

		name string
	}

	stringReader struct {
		*paramReader

		stripWhitespaces bool
		notEmpty         bool
	}

	numRangeReader struct {
		*stringReader

		from int
		to   int
	}
)

func newParamReader(printer printer, scanner scanner, name string) *paramReader {
	return &paramReader{
		printer: printer,
		scanner: scanner,

		name: name,
	}
}

func (r *paramReader) String() *stringReader {
	return &stringReader{
		paramReader:      r,
		stripWhitespaces: true,
		notEmpty:         true,
	}
}

func (r *paramReader) NumRange(from, to int) *numRangeReader {
	return &numRangeReader{
		stringReader: r.String(),

		from: from,
		to:   to,
	}
}

func (r *paramReader) Month() *numRangeReader {
	return r.NumRange(int(time.January), int(time.December))
}

func (r *paramReader) Day() *numRangeReader {
	return r.NumRange(1, 31)
}

func (r *stringReader) StripWhitespaces(s bool) *stringReader {
	r.stripWhitespaces = s
	return r
}

func (r *stringReader) NotEmpty(ne bool) *stringReader {
	r.notEmpty = ne
	return r
}

func (r *stringReader) Read() (string, error) {
	r.printer.Printf("Enter the %s: ", r.name)

	res, err := r.scanner.Readln()
	if err != nil {
		return "", errors.Errorf("failed to read a %s", r.name)
	}

	if r.stripWhitespaces {
		res = stripWhiteSpace(res)
	}

	if r.notEmpty && res == "" {
		return "", errors.Errorf("%s should not be empty", r.name)
	}

	return res, nil
}

func (r *numRangeReader) Read() (int, error) {
	strMonth, err := r.stringReader.Read()
	if err != nil {
		return 0, err
	}

	intMonth, err := strconv.Atoi(strMonth)
	if err != nil {
		return 0, errors.Wrapf(err, "should be a number, got %q", strMonth)
	}

	if intMonth < r.from || intMonth > r.to {
		return 0, errors.Errorf("%s should be a number between %d and %d inclusively, got %d", r.name, r.from, r.to, intMonth)
	}

	return intMonth, nil
}

// -- Helpers --

func stripWhiteSpace(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}
