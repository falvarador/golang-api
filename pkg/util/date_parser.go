package util

import (
	"fmt"
	"time"
)

// ParseTimeFromString attempts to parse a date string to time.Time using RFC3339Nano.
// If the parsing fails, it returns an error with the original string for debugging.
func ParseTimeFromString(dateString string) (time.Time, error) {
	parsedTime, err := time.Parse(time.RFC3339Nano, dateString)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse time string '%s' to RFC3339Nano: %w", dateString, err)
	}
	return parsedTime, nil
}

// FormatTimeToString formats a time.Time to a string using RFC3339Nano.
func FormatTimeToString(t time.Time) string {
	return t.Format(time.RFC3339Nano)
}
