package hello_test

import (
	"testing"
	"bytes"
	"github.com/mik135/hello"
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