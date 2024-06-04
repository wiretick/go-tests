package racer

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

var ErrTimedOut = errors.New("Request timed out")

func ping(url string) chan struct{} {
	res := make(chan struct{})
	go func() {
		http.Get(url)
		close(res)
	}()
	return res
}

func Racer(one, two string, timeout time.Duration) (string, error) {
	select {
	case <-ping(one):
		return one, nil
	case <-ping(two):
		return two, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out")
	}
}
