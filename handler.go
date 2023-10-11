package metrica

import (
	"encoding/json"
	"net/http"
	"sync"
	"time"

	"log/slog"
)

type countResponse struct {
	Count int64
}

// Handler returns a http.Handler for new endpoints.
func Handler(c *AtomicCounter, fs *FileStorage) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/count", func(w http.ResponseWriter, r *http.Request) {
		counter(w, r, c)
	})
	mux.HandleFunc("/countfs", func(w http.ResponseWriter, r *http.Request) {
		counterFs(w, r, fs)
	})
	return mux
}

func counter(w http.ResponseWriter, _ *http.Request, c *AtomicCounter) {
	c.Inc(1)
	send(w, http.StatusOK, countResponse{Count: c.Value()})
}

func counterFs(w http.ResponseWriter, _ *http.Request, fs *FileStorage) {
	var wg sync.WaitGroup
	var counter int64

	c := Counter(time.Now())

	for i := 0; i < 1; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := fs.Write(c)
			if err != nil {
				send(w, http.StatusInternalServerError, err)
				return
			}

			counters, err := fs.Read()
			if err != nil {
				send(w, http.StatusInternalServerError, err)
				return
			}
			counter = counters.Count60sec()
		}()
	}
	wg.Wait()
	send(w, http.StatusOK, countResponse{Count: counter})
}

func send(w http.ResponseWriter, statusCode int, body interface{}) {
	const jsonContentType = "application/json; charset=utf-8"

	w.Header().Set("Content-Type", jsonContentType)
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(body); err != nil {
		slog.Error("Unable to encode body as JSON", "error", err)
	}
}
