package context

import (
	"context"
	"fmt"
	"net/http"
)

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// ctx := r.Context()

		// data := make(chan string, 1)

		// go func() {
		// 	data <- store.Fetch()
		// }()

		// select {
		// case d := <-data:
		// 	fmt.Fprintf(w, d)
		// case <-ctx.Done():
		// 	store.Cancel()
		// }
		data, _ := store.Fetch(r.Context())
		fmt.Fprintf(w, data)
	}
}

type Store interface {
	Fetch(ctx context.Context) (string, error)
}
