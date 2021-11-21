package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/pldi-submission-448/fgg/internal/fg"
	"github.com/pldi-submission-448/fgg/internal/fgg"
	"github.com/pldi-submission-448/fgg/internal/parser"
)

func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		fmt.Fprintln(os.Stderr, "not enough arguments (expected FGG file path)")
		os.Exit(1)
	}
	b, err := ioutil.ReadFile(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}
	fgAdaptor := new(parser.FGAdaptor)
	fgProg := fgAdaptor.Parse(false, string(b))
	fggProg, err := fgg.FromFG(fgProg.(fg.FGProgram))
	if err != nil {
		log.Fatalf("cannot convert from FG program: %v", err)
	}
	fmt.Println(fggProg.String())
}
