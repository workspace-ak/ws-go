package mock

import (
	"fmt"
	"io"
	"time"
)

const finalWord = "Go!"

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func Countdown(out io.Writer, n int, s Sleeper) {
	for i := n; i > 0; i-- {
		fmt.Fprintln(out, i)
		s.Sleep()

	}
	fmt.Fprint(out, finalWord)

}
