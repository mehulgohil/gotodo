package main

import (
	"flag"
	"fmt"
)

func main() {
	name := flag.String("name", "", "The name to greet")
	flag.Parse()

	fmt.Println("Hello,", *name)
}
