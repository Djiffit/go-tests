package context

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type StubStore struct {
	response string
}

func (s *StubStore) Fetch() string {
	return s.response
}

func (s *StubStore) Cancel() {

}

type SpyStore struct {
	response string
	t        *testing.T
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	go func() {
		var result string
		for _, c := range s.response {
			select {
			case <-ctx.Done():
				s.t.Log("spy store got cancelled")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		}
		data <- result
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}
}

// func (s *SpyStore) Cancel() {
// 	// s.cancelled = true
// }

func TestHandler(t *testing.T) {
	// t.Run("normally working", func(t *testing.T) {
	// 	data := "hello, world"
	// 	store := &SpyStore{data, t}
	// 	svr := Server(store)

	// 	request := httptest.NewRequest(http.MethodGet, "/", nil)
	// 	response := httptest.NewRecorder()

	// 	svr.ServeHTTP(response, request)

	// 	if response.Body.String() != data {
	// 		t.Errorf(`got %s, want %s`, response.Body.String(), data)
	// 	}

	// 	if store.cancelled {
	// 		t.Error("Request was not supposed to be cancelled")
	// 	}
	// })

	// t.Run("works with cancel properly", func(t *testing.T) {
	// 	data := "hello, world"
	// 	store := &SpyStore{data, t}
	// 	svr := Server(store)

	// 	request := httptest.NewRequest(http.MethodGet, "/", nil)
	// 	cancelCtx, cancel := context.WithCancel(request.Context())
	// 	time.AfterFunc(5*time.Millisecond, cancel)
	// 	request = request.WithContext(cancelCtx)

	// 	response := httptest.NewRecorder()
	// 	svr.ServeHTTP(response, request)

	// 	if !store.cancelled {
	// 		t.Error("store was supposed to be cancelled")
	// 	}
	// })

	t.Run("returns data from store", func(t *testing.T) {
		data := "hello, world"
		store := &SpyStore{data, t}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf(`got "%s", want "%s"`, response.Body.String(), data)
		}
	})

}
