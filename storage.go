// Package metrica (storage.go) implements a version of a storage using the filesystem
package metrica

import (
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"time"
)

// Counter is an alias for time.Time
type Counter time.Time

// Counters is a list of Counters
type Counters []Counter

// FileStorage is a struct that implements the Storage methods
type FileStorage struct {
	mu       *sync.Mutex
	fileName string
}

// NewFileStorage initializes a new FileStorage
func NewFileStorage(mu *sync.Mutex, fileName string) *FileStorage {
	return &FileStorage{
		mu:       mu,
		fileName: fileName,
	}
}

// Write save a Counter in the file
func (m *FileStorage) Write(c Counter) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	f, err := os.OpenFile(m.fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("error opening file: %v", err)
	}

	_, err = f.WriteString(fmt.Sprintf("%v\n", c.Format()))
	if err != nil {
		return fmt.Errorf("error writing file: %v", err)
	}
	return nil
}

// Read returns a list of Counters
func (m *FileStorage) Read() (Counters, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	f, err := os.Open(m.fileName)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}

	defer func() {
		_ = f.Close()
	}()

	fileData, err := io.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	datetimes := strings.Split(string(fileData), "\n")
	if datetimes[len(datetimes)-1] == "" {
		datetimes = datetimes[:len(datetimes)-1]
	}

	return parseDatetimesToCounters(datetimes)
}

// Format returns the datetime in RFC3339Nano format
func (c Counter) Format() string {
	return time.Time(c).Format(time.RFC3339Nano)
}

// Count60sec returns the number of datetimes in the last 60 seconds
func (cs Counters) Count60sec() int64 {
	currentTime := time.Now()
	var count int64

	lowerLimit := currentTime.Add(-60 * time.Second)
	for _, c := range cs {
		t := time.Time(c)

		if t.After(lowerLimit) && t.Before(currentTime) {
			count++
		}
	}

	return count
}

func parseDatetimesToCounters(datetimes []string) (Counters, error) {
	var counters []Counter
	for _, datetime := range datetimes {
		t, err := time.Parse(time.RFC3339Nano, datetime)
		if err != nil {
			return nil, fmt.Errorf("error parsing datetime: %v", err)
		}
		counters = append(counters, Counter(t))
	}
	return counters, nil
}
