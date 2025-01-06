package main

import (
	"flag"
	"fmt"

	"github.com/fatjyc/hello-go/pkg/hello"
)

var lang = flag.String("lang", "en", "language")
var name = flag.String("name", "World", "name")

func main() {
	flag.Parse()
	greeter := hello.NewGreeter(*lang, *name)
	fmt.Println(greeter.Greet())
}
