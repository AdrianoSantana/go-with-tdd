package cumprimenta

import (
	"fmt"
	"io"
)

func Cumprimenta(w io.Writer, palavra string) {
	fmt.Fprintf(w, "Olá, %s", palavra)
}
