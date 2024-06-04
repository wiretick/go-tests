package racer

import (
	"net/http"
	"time"
)

func measureResTime(url string) time.Duration {
	startOne := time.Now()
	http.Get(url)
	return time.Since(startOne)
}

func Racer(one, two string) string {
	tookOne := measureResTime(one)
	tookTwo := measureResTime(two)

	if tookOne < tookTwo {
		return one
	}

	return two
}
