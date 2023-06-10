package printer

import (
	"fmt"
	"os"
)

type printer struct {
	out    *os.File
	errOut *os.File
}

func New(out *os.File, errOut *os.File) *printer {
	return &printer{
		out:    out,
		errOut: errOut,
	}
}

func (p *printer) Printf(format string, a ...any) {
	fmt.Fprintf(p.out, format, a...)
}

func (p *printer) EPrinteln(a ...any) {
	fmt.Fprintln(p.errOut, a...)
}
