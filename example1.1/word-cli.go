package main

import (
	"flag"
	"fmt"
)

func main() {
	// define a flag with it's type with default and help text
	wordPtr := flag.String("word", "some default value for word", "some description of word")

	// read the flags provided
	flag.Parse()

	// print/format the pointer value
	fmt.Printf("'%s' is the word\n", *wordPtr)
}
