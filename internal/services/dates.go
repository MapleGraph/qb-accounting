package services

import (
	"fmt"
	"time"
)

func parseDateTime(s string) (time.Time, error) {
	if s == "" {
		return time.Time{}, fmt.Errorf("empty date string")
	}
	layouts := []string{
		time.RFC3339Nano,
		time.RFC3339,
		"2006-01-02T15:04:05Z07:00",
		"2006-01-02",
	}
	var last error
	for _, layout := range layouts {
		t, err := time.Parse(layout, s)
		if err == nil {
			return t, nil
		}
		last = err
	}
	return time.Time{}, fmt.Errorf("invalid date %q: %w", s, last)
}
