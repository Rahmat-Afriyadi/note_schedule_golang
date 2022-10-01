package dto

import (
	"note_scheduler/entity"
	// "note_scheduler/entity"
)

type ScheduleInsertDTO struct {
	Judul      string            `json:"judul" validate:"required" form:"judul" bind:"required"`
	StartAt    string            `json:"start_at" validate:"required" form:"start_at" bind:"required"`
	EndAt      string            `json:"end_at" validate:"required" form:"end_at" bind:"required"`
	Activities []entity.Activity `json:"activities" form:"activities" bind:"required"`
}

type ScheduleUpdateDTO struct {
	Id      uint64 `json:"id" validate:"required" form:"id" bind:"required"`
	Judul   string `json:"judul" validate:"required" form:"judul" bind:"required"`
	StartAt string `json:"start_at" validate:"required" form:"start_at" bind:"required"`
	EndAt   string `json:"end_at" validate:"required" form:"end_at" bind:"required"`
}

type OnGoingSchedule struct {
	From string `json:"from" validate:"required" form:"from" bind:"required"`
	To   string `json:"to" validate:"required" form:"to" bind:"required"`
}
