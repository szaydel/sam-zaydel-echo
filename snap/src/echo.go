package main

import (
	"flag"
	"fmt"
	"strings"
)

var noNewLineFlag bool
var tokenSepFlag bool

func main() {
	flag.BoolVar(&noNewLineFlag, "n", false, "Do not output newline")
	flag.BoolVar(&tokenSepFlag, "s", false, "Do not separate arguments with spaces")
	flag.Parse()
	var sep = " "
	if tokenSepFlag {
		sep = ""
	}
	if noNewLineFlag {
		joined := strings.Join(flag.Args(), sep)
		fmt.Printf("%s", joined)
		return
	} else {
		joined := strings.Join(flag.Args(), sep)
		fmt.Printf("%s\n", joined)
	}
}
