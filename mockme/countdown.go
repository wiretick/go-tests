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

func Countdown(w io.Writer) {
	for i := countdown; i > 0; i-- {
		fmt.Fprintf(w, "%d ", i)
		time.Sleep(1 * time.Second)
	}
	fmt.Fprintf(w, ending)
}

func main() {
	Countdown(os.Stdout)
}
