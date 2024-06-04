package racer

import (
	"fmt"
	"net/http"
	"time"
)

func ping(url string) chan struct{} {
	res := make(chan struct{})
	go func() {
		http.Get(url)
		close(res)
	}()
	return res
}

func Racer(one, two string) (string, error) {
	select {
	case <-ping(one):
		return one, nil
	case <-ping(two):
		return two, nil
	case <-time.After(10 * time.Second):
		return "", fmt.Errorf("timed out")
	}
}
