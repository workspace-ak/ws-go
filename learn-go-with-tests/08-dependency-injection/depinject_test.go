package depinject

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buf := bytes.Buffer{}
	Greet(&buf, "Chris")
	res := buf.String()
	expect := "Hello, Chris"
	if res != expect {
		t.Errorf("got %q, want %q", res, expect)
	}
}
