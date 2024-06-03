package mockme

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	ending    = "GO!"
	countdown = 3
)

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

func Countdown(w io.Writer, sleeper Sleeper) {
	for i := countdown; i > 0; i-- {
		fmt.Fprintf(w, "%d ", i)
		sleeper.Sleep()
	}
	fmt.Fprintf(w, ending)
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
