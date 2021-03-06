package fg

import (
	"strings"
)

// Pre: len(elems) > 1
// Pre: elems[:len(elems)-1] -- type/meth Decls; elems[len(elems)-1] -- "main" func body expression
func MakeFgProgram(elems ...string) string {
	if len(elems) == 0 {
		panic("Bad empty args: must supply at least body expression for \"main\"")
	}
	var b strings.Builder
	b.WriteString("package main;\n")
	for _, v := range elems[:len(elems)-1] {
		b.WriteString(v)
		b.WriteString(";\n")
	}
	b.WriteString("func main() { _ = " + elems[len(elems)-1] + " }")
	return b.String()
}
