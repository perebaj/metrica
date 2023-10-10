package metrica

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
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

func assert(t *testing.T, want interface{}, got interface{}) {
	if want != got {
		t.Errorf("expected %v; got %v", want, got)
	}
}
