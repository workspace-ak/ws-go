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

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

// Countdown from n with 1 second interval
func Countdown(out io.Writer, n int, s Sleeper) {
	for i := n; i > 0; i-- {
		fmt.Fprintln(out, i)
		s.Sleep()

	}
	fmt.Fprint(out, finalWord)

}
