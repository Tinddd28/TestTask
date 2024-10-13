package utils

import (
	"time"
)

func ParseDate(date string) (time.Time, error) {
	layout := "02.01.2006"

	parsedDate, err := time.Parse(layout, date)
	if err != nil {
		return time.Time{}, err
	}

	return parsedDate, nil
}
