package metrica

import (
	"encoding/json"
	"net/http"

	"log/slog"
)

type countResponse struct {
	Count int
}

// Handler returns a http.Handler for new endpoints.
func Handler(c *AtomicCounter) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/count", func(w http.ResponseWriter, r *http.Request) {
		counter(w, r, c)
	})
	return mux
}

func counter(w http.ResponseWriter, _ *http.Request, c *AtomicCounter) {
	c.Inc(1)
	send(w, http.StatusOK, countResponse{Count: int(c.Value())})
}

func send(w http.ResponseWriter, statusCode int, body interface{}) {
	const jsonContentType = "application/json; charset=utf-8"

	w.Header().Set("Content-Type", jsonContentType)
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(body); err != nil {
		slog.Error("Unable to encode body as JSON", "error", err)
	}
}
