package racer

import (
	"errors"
	"net/http"
	"time"
)

var ErrTimedOut = errors.New("Request timed out")
var defaultTimeout = 10 * time.Second

func ping(url string) chan struct{} {
	res := make(chan struct{})
	go func() {
		http.Get(url)
		close(res)
	}()
	return res
}

func Racer(one, two string) (string, error) {
	return ConfigurableRacer(one, two, defaultTimeout)
}

func ConfigurableRacer(one, two string, timeout time.Duration) (string, error) {
	select {
	case <-ping(one):
		return one, nil
	case <-ping(two):
		return two, nil
	case <-time.After(timeout):
		return "", ErrTimedOut
	}
}
