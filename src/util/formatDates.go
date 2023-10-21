package util

import (
	"errors"
	"time"
)

func FormatTime(dateString string) (string, error) {
	parsedTime, err := time.Parse("2006-01-02T03:04:05Z", dateString)
	if err != nil {
		return "", errors.New("503")
	}
	formattedTime := parsedTime.Format(time.RFC1123)

	return formattedTime, nil
}
