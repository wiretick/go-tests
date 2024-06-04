package racer

import (
	"net/http"
)

func Racer(one, two string) string {
	resOne := make(chan struct{})
	resTwo := make(chan struct{})

	go func() {
		http.Get(one)
		close(resOne)
	}()

	go func() {
		http.Get(two)
		close(resTwo)
	}()

	select {
	case <-resOne:
		return one
	case <-resTwo:
		return two
	}
}
