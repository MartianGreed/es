package es

import (
	"log"
	"time"
)

func StringifyDate(t time.Time) string {
	return t.Format(time.RFC3339)
}

func ParseDate(t string) time.Time {
	parsedTime, err := time.Parse(time.RFC3339, t)
	if err != nil {
		log.Fatalf("%+v\n", err)
	}

	return parsedTime
}
