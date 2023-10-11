package metrica

import (
	"os"
	"sync"
	"testing"
	"time"
)

func TestWriteRead(t *testing.T) {
	c := Counter(time.Now())

	f, err := os.CreateTemp("", "test")
	if err != nil {
		t.Fatalf("error creating temp file: %v", err)
	}
	defer func() {
		_ = os.Remove(f.Name())
	}()

	fs := NewFileStorage(&sync.Mutex{}, f.Name())

	if err := fs.Write(c); err != nil {
		t.Fatalf("error writing to file: %v", err)
	}

	got, err := fs.Read()
	if err != nil {
		t.Fatalf("error reading file: %v", err)
	}

	if len(got) == 1 {
		assert(t, c.Format(), got[0].Format())
	} else {
		t.Fatalf("expected 1; got %v", len(got))
	}
}

func TestCountersCount60sec(t *testing.T) {
	currentTime := time.Now()
	counters := Counters{
		Counter(currentTime.Add(-time.Second * 10)),
		Counter(currentTime.Add(-time.Second * 10)),
		Counter(currentTime.Add(-time.Second * 20)),
		Counter(currentTime.Add(-time.Second * 30)),
		Counter(currentTime.Add(-time.Second * 40)),
		Counter(currentTime.Add(-time.Second * 50)),
		Counter(currentTime.Add(-time.Second * 60)),
	}
	want := int64(6)
	got := counters.Count60sec()
	if got != want {
		t.Fatalf("expected %v; got %v", want, got)
	}

}
