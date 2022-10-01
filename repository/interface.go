package repository

import (
	"note_scheduler/entity"
)

type ActivityRepository interface {
	All() []entity.Activity
	FindById(id uint64) entity.Activity
	Insert(e entity.Activity) entity.Activity
	Update(e entity.Activity) entity.Activity
	Delete(id uint64) error
	Where(e entity.Activity) []entity.Activity
	
}

type ScheduleRepository interface {
	All() []entity.Schedule
	FindById(id uint64) entity.Schedule
	Insert(e entity.Schedule) entity.Schedule
	Update(e entity.Schedule) entity.Schedule
	Delete(id uint64)
	OnGoingSchedule(from string, to string) []entity.Schedule
	ByJudul(judul string) []entity.Schedule
}
