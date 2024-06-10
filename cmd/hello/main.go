package main

import (
	"github.com/mik135/hello"
	"os"
)

func main() {
	hello.PrintTo(os.Stdout)
	hello.GreetUser(os.Stdin, os.Stdout)
}