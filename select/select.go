package racer

import (
	"fmt"
	"net/http"
	"time"
)

// Racer races two strings
func Racer(a, b string, timeout time.Duration) (winner string, err error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out while waiting for %s and %s", a, b)

	}

}

func ping(url string) chan bool {
	ch := make(chan bool)

	go func() {
		http.Get(url)
		ch <- true
	}()

	return ch
}

func measureTime(req string) time.Duration {
	start := time.Now()
	http.Get(req)
	return time.Since(start)
}
