package metrica

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

// Counter is a struct that holds the datetime of the request
type Counter struct {
	Datetime string
}

// Write writes the datetime to the file
func Write(f *os.File, c Counter) error {
	_, err := f.WriteString(fmt.Sprintf("%v\n", c.Datetime))
	if err != nil {
		return fmt.Errorf("error writing file: %v", err)
	}
	return nil
}

// Read reads the datetime from the file
func Read(filename string) ([]time.Time, error) {
	f, err := os.Open(filename)
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

	var parsedDatetime []time.Time
	for _, datetime := range datetimes {
		t, err := time.Parse(time.RFC3339Nano, datetime)
		if err != nil {
			return nil, fmt.Errorf("error parsing datetime: %v", err)
		}
		parsedDatetime = append(parsedDatetime, t)

	}
	return parsedDatetime, nil
}
