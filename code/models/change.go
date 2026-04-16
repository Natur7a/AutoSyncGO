package models

import "time"

type Change struct {
	ID         int
	FilePath   string
	ChangeType string
	ChangedAt  time.Time
}
