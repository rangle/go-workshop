package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Println("Hello, Rangle")
	wordPtr := flag.String("word", "some default value for word", "some description of word")
	flag.Parse()
	fmt.Println(*wordPtr)
}
