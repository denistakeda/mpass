package scanner

import (
	"bufio"
	"os"
)

type scanner struct {
	in *os.File
}

func New(in *os.File) *scanner {
	return &scanner{in: in}
}

func (s *scanner) Readln() (string, error) {
	// FIXME: something is wrong here
	sc := bufio.NewScanner(s.in)
	sc.Scan()
	return sc.Text(), sc.Err()
}
