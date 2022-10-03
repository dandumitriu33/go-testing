package contextdemo

import (
	"context"
	"fmt"
	"net/http"
)

type Store interface {
	Fetch(ctx context.Context) (string, error)
}

// type Store interface {
// 	Fetch() string
// 	Cancel()
// }

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, _ := store.Fetch(r.Context())
		fmt.Fprint(w, data)
	}
}

// func Server(store Store) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		ctx := r.Context()

// 		data := make(chan string, 1)

// 		go func() {
// 			data <- store.Fetch()
// 		}()

// 		select {
// 		case d := <-data:
// 			fmt.Fprint(w, d)
// 		case <-ctx.Done():
// 			store.Cancel()
// 		}
// 	}
// }
