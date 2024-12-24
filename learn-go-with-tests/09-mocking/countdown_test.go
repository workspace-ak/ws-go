package mock

import (
	"bytes"
	"os"
	"reflect"
	"testing"
	"time"
)

func TestCountdown(t *testing.T) {
	t.Run("prints 3 to go!", func(t *testing.T) {
		buf := &bytes.Buffer{}
		n := 3
		// spySleeper := &SpySleeper{}
		Countdown(buf, n, &SpyCountDownOperations{})
		res := buf.String()
		expect := `3
2
1
Go!`
		if res != expect {
			t.Errorf("got %q, want %q", res, expect)
		}
	})
	t.Run("sleep before every print", func(t *testing.T) {
		n := 3
		spySleepPrinter := &SpyCountDownOperations{}
		Countdown(spySleepPrinter, n, spySleepPrinter)
		res := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}
		if !reflect.DeepEqual(res, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", res, spySleepPrinter.Calls)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second
	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()
	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}

func ExampleCountdown() {
	n := 3
	sleeper := &ConfigurableSleeper{time.Duration(n) * time.Second, time.Sleep}
	Countdown(os.Stdout, n, sleeper)
	// Output: 3
	// 2
	// 1
	// Go!
}

const sleep = "sleep"
const write = "write"

type SpyCountDownOperations struct {
	Calls []string
}

func (s *SpyCountDownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountDownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}
