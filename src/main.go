package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")

	var N = flag.Int64("N", 0, "type your N")
	var M = flag.Int64("M", 0, "type your M")
	flag.Parse()

	fmt.Printf("Pemain = %d, Dadu = %d", *N, *M)
}
