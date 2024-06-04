package racer

import (
	"net/http"
	"time"
)

func Racer(one, two string) string {
	startOne := time.Now()
	http.Get(one)
	tookOne := time.Since(startOne)

	startTwo := time.Now()
	http.Get(two)
	tookTwo := time.Since(startTwo)

	if tookOne < tookTwo {
		return one
	}

	return two
}
