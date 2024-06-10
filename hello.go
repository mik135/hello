package hello

import (
	"bufio"
	"fmt"
	"io"
)

func PrintTo(w io.Writer) {
	fmt.Fprintln(w, "Hello, world")
}

func GreetUser(stdin io.Reader, stdout io.Writer) {
	name := "you"
	fmt.Fprintln(stdout, "What is your name?")
	input := bufio.NewScanner(stdin)
	if input.Scan() {
		name = input.Text()
	}
	fmt.Fprintf(stdout, "Hello, %s.\n", name)
}