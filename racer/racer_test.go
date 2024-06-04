package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func timedServer(duration time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(duration * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))
}

func TestRacer(t *testing.T) {
	slowServer := timedServer(10)
	defer slowServer.Close()

	fastServer := timedServer(0)
	defer fastServer.Close()

	slow := slowServer.URL
	fast := fastServer.URL

	got := Racer(slow, fast)
	want := fast

	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}
