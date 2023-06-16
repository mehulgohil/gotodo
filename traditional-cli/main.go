package main

import (
	"flag"
	"fmt"
)

func main() {
	name := flag.String("name", "Mag", "The name to greet")
	flag.Parse()

	args := flag.Args()

	fmt.Println("Hello,", *name)
	fmt.Println("Args:", args)
}
