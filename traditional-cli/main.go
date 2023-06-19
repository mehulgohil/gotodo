package main

import (
	"flag"
	"fmt"
)

func main() {
	name := flag.String("name", "mehul", "The name to greet")
	flag.Parse()

	fmt.Println("Hello,", *name)
}
