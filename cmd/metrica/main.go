// Packagem main gather all the packages and run the server
package main

import (
	"log/slog"
	"net/http"

	"github.com/perebaj/metrica"
)

func main() {
	c := metrica.NewAtomicCounter()

	mux := metrica.Handler(c)
	slog.Info("Starting server", "port", 8080)
	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		panic(err)
	}
}
