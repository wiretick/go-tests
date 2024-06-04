package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func timedServer(duration time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(duration)
		w.WriteHeader(http.StatusOK)
	}))
}

func TestRacer(t *testing.T) {
	t.Run("find fastest server", func(t *testing.T) {
		slowServer := timedServer(10 * time.Millisecond)
		defer slowServer.Close()

		fastServer := timedServer(0 * time.Millisecond)
		defer fastServer.Close()

		slow := slowServer.URL
		fast := fastServer.URL

		got, _ := Racer(slow, fast)
		want := fast

		if want != got {
			t.Errorf("want %q, got %q", want, got)
		}
	})

	t.Run("error if times out", func(t *testing.T) {
		slowServer := timedServer(11 * time.Second)
		defer slowServer.Close()

		fastServer := timedServer(12 * time.Second)
		defer fastServer.Close()

		_, err := Racer(slowServer.URL, fastServer.URL)

		if err == nil {
			t.Errorf("expected error, but got none")
		}
	})
}
