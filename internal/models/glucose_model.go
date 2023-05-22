package models

import "time"

type GlucoseLevel struct {
	Type  string    `json:"type"`
	Level float32   `json:"level"`
	Date  time.Time `json:"date"`
}
