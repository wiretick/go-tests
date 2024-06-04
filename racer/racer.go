package racer

import (
	"net/http"
)

func ping(url string) chan struct{} {
	res := make(chan struct{})
	go func() {
		http.Get(url)
		close(res)
	}()
	return res
}

func Racer(one, two string) string {
	select {
	case <-ping(one):
		return one
	case <-ping(two):
		return two
	}
}
