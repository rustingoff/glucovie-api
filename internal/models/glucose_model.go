package models

import "time"

type GlucoseLevel struct {
	Type  string    `json:"type"`
	Level string    `json:"level"`
	Date  time.Time `json:"date"`
}

type GlucoseResponse struct {
	Type  string `json:"type"`
	Level string `json:"level"`
	Day   string `json:"day"`
}
