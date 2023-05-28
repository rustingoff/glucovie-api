package models

import "time"

type GlucoseLevel struct {
	Type   string    `json:"type"`
	Level  float32   `json:"level"`
	Date   time.Time `json:"date"`
	UserID string    `json:"user_id"`
}

type GlucoseResponse struct {
	Type  string  `json:"type"`
	Level float32 `json:"level"`
	Day   int32   `json:"day"`
}
