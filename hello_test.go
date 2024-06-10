package hello_test

import (
	"testing"
	"bytes"
	"github.com/mik135/hello"
	"testing/iotest"
	"errors"
)

func TestPrintPrintsHelloMessageToTerminal(t *testing.T) {
	t.Parallel()
	buf := new(bytes.Buffer)
	hello.PrintTo(buf)
	want := "Hello, world\n"
	got := buf.String()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestGreetGreetsUser(t *testing.T) {
	t.Parallel()
	input := iotest.ErrReader(errors.New("bad reader"))
	output := new(bytes.Buffer)
	hello.GreetUser(input, output)
	want := "What is your name?\nHello, you.\n"
	got := output.String()

	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}

}