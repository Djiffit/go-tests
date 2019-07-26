package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("returns correct url", func(t *testing.T) {
		slow := delayedServer(20 * time.Millisecond)
		fast := delayedServer(0)

		defer slow.Close()
		defer fast.Close()

		want := fast.URL
		got, _ := Racer(slow.URL, fast.URL, 1*time.Second)

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("timeouts if request > 10 seconds", func(t *testing.T) {
		slow := delayedServer(100 * time.Millisecond)

		defer slow.Close()

		_, err := Racer(slow.URL, slow.URL, 50*time.Millisecond)

		if err == nil {
			t.Error("Expected an error didn't get one")
		}
	})
}

func delayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
