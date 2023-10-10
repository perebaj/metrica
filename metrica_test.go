package metrica

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
)

func TestHandlerCount(t *testing.T) {
	req := httptest.NewRequest("GET", "/count", nil)
	w := httptest.NewRecorder()
	c := AtomicCounter{}
	mux := Handler(&c)
	mux.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("expected status OK; got %v", w.Code)
	}

	var got countResponse
	if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
		t.Errorf("unable to decode body: %v", err)
	}

	assert(t, 1, got.Count)
}

func TestHandlerCount_Multiple(t *testing.T) {
	c := NewAtomicCounter()
	mux := Handler(c)

	for i := 0; i < 100; i++ {
		req := httptest.NewRequest("GET", "/count", nil)
		w := httptest.NewRecorder()
		mux.ServeHTTP(w, req)
		if w.Code != http.StatusOK {
			t.Errorf("expected status OK; got %v", w.Code)
		}

		var gotRes countResponse
		if err := json.NewDecoder(w.Body).Decode(&gotRes); err != nil {
			t.Errorf("unable to decode body: %v", err)
		}

		wantCount := i + 1
		assert(t, wantCount, gotRes.Count)
	}
}

func TestHandlerCount_Concurrent(t *testing.T) {
	c := NewAtomicCounter()
	mux := Handler(c)
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			req := httptest.NewRequest("GET", "/count", nil)
			w := httptest.NewRecorder()
			mux.ServeHTTP(w, req)
			if w.Code != http.StatusOK {
				t.Errorf("expected status OK; got %v", w.Code)
			}
			defer wg.Done()
		}()
	}
	wg.Wait()

	assert(t, int64(1000), c.Value())
}

func assert(t *testing.T, want interface{}, got interface{}) {
	if want != got {
		t.Errorf("expected %v; got %v", want, got)
	}
}
