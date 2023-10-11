// Packagem main gather all the packages and run the server
package main

import (
	"log/slog"
	"net/http"
	"sync"

	"github.com/perebaj/metrica"
)

func main() {
	c := metrica.NewAtomicCounter()
	fs := metrica.NewFileStorage(&sync.Mutex{}, "counters.txt")
	mux := metrica.Handler(c, fs)

	slog.Info("Starting server", "port", 8080)
	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		panic(err)
	}
}
