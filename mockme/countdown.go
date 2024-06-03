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

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func Countdown(w io.Writer, sleep Sleeper) {
	for i := countdown; i > 0; i-- {
		fmt.Fprintf(w, "%d ", i)
	}
	for i := countdown; i > 0; i-- {
		sleep.Sleep()
	}
	fmt.Fprintf(w, ending)
}

func main() {
	sleep := &DefaultSleeper{}
	Countdown(os.Stdout, sleep)
}
