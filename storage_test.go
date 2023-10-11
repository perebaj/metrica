package metrica

import (
	"os"
	"testing"
	"time"
)

func TestWriteRead(t *testing.T) {
	c := Counter{
		Datetime: time.Now().Format(time.RFC3339Nano),
	}

	f, err := os.CreateTemp("", "test")
	if err != nil {
		t.Fatalf("error creating temp file: %v", err)
	}
	defer os.Remove(f.Name())

	if err := Write(f, c); err != nil {
		t.Fatalf("error writing to file: %v", err)
	}

	got, err := Read(f.Name())
	if err != nil {
		t.Fatalf("error reading file: %v", err)
	}

	if len(got) == 1 {
		assert(t, got[0].Format(time.RFC3339Nano), c.Datetime)
	} else {
		t.Fatalf("expected 1; got %v", len(got))
	}
}
