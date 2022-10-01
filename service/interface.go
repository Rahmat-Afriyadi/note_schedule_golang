package service

import (
	"note_scheduler/dto"
	"note_scheduler/entity"

)


type ScheduleService interface {
	All() []entity.Schedule
	FindById(id uint64) entity.Schedule
	Create(dto dto.ScheduleInsertDTO) (entity.Schedule, error)
	Update(dto dto.ScheduleUpdateDTO) (entity.Schedule, error)
	Delete(id uint64)
	OnGoingSchedule(from string, to string) []entity.Schedule
	ByJudul(judul string) []entity.Schedule
}

type ActivityService interface {
	All() []entity.Activity
	FindById(id uint64) entity.Activity
	Where(e entity.Activity) []entity.Activity
	Create(dto dto.ActivityInsertDTO) (entity.Activity, error)
	Update(dto dto.ActivityUpdateDTO) (entity.Schedule, error)
	Delete(id uint64) entity.Schedule
}
