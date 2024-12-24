package mock

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	buf := &bytes.Buffer{}
	n := 3
	spySleeper := &SpySleeper{}
	Countdown(buf, n, spySleeper)
	res := buf.String()
	expect := `3
2
1
Go!`
	if res != expect {
		t.Errorf("got %q, want %q", res, expect)
	}
	if spySleeper.Calls != n {
		t.Errorf("not enough calls to sleeper, want %d got %d", n, spySleeper.Calls)
	}

}
