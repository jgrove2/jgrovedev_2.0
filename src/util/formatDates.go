package util

import (
	"log"
	"time"
)

func FormatTime(dateString string) string {
	parsedTime, err := time.Parse("2006-01-02T03:04:05Z", dateString);
	if err != nil {
		log.Fatalf("Error parsing string %v", err)
	}
	formattedTime := parsedTime.Format(time.RFC1123)

	return formattedTime
}