package entity

import (
	"time"
)

type Schedule struct {
	Id      uint64    `gorm:"primaryKey:autoIncreament" json:"id"`
	StartAt time.Time `json:"start_at"`
	EndAt   time.Time `json:"end_at"`
	Judul   string    `json:"judul"`

	Activities []Activity `json:"activities,omitempty"`
}
