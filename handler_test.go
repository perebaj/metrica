package metrica

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"sync"
	"testing"
)

func TestHandlerCountFS_Sequential(t *testing.T) {
	f, err := os.Create("test")
	if err != nil {
		t.Fatalf("error creating temp file: %v", err)
	}
	defer func() {
		_ = os.Remove(f.Name())
	}()

	fs := NewFileStorage(&sync.Mutex{}, f.Name())
	mux := Handler(fs)

	for i := 0; i < 100; i++ {
		req := httptest.NewRequest("GET", "/countfs", nil)
		w := httptest.NewRecorder()
		mux.ServeHTTP(w, req)
		if w.Code != http.StatusOK {
			t.Errorf("expected status OK; got %v", w.Code)
		}

		var gotRes countResponse
		if err := json.NewDecoder(w.Body).Decode(&gotRes); err != nil {
			t.Errorf("unable to decode body: %v", err)
		}

		wantCount := int64(i + 1)
		assert(t, wantCount, gotRes.Count)
	}
}

func TestHandlerCountFS_Concurrent(t *testing.T) {
	f, err := os.Create("test")
	if err != nil {
		t.Fatalf("error creating temp file: %v", err)
	}
	defer func() {
		_ = os.Remove(f.Name())
	}()

	fs := NewFileStorage(&sync.Mutex{}, f.Name())

	mux := Handler(fs)
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			req := httptest.NewRequest("GET", "/countfs", nil)
			w := httptest.NewRecorder()
			mux.ServeHTTP(w, req)
			if w.Code != http.StatusOK {
				t.Errorf("expected status OK; got %v", w.Code)
			}
			defer wg.Done()
		}()
	}
	wg.Wait()

	got, err := fs.Read()
	if err != nil {
		t.Fatalf("error reading file: %v", err)
	}
	assert(t, int64(100), got.Count60sec())
}

func assert(t *testing.T, want interface{}, got interface{}) {
	if want != got {
		t.Errorf("expected %v; got %v", want, got)
	}
}
