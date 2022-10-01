package service

import (
	"note_scheduler/dto"
	"note_scheduler/entity"
	"note_scheduler/repository"
	"time"

	"github.com/go-playground/validator/v10"
)

type ScheduleServiceImpl struct {
	repository.ScheduleRepository
	*validator.Validate
}

func NewScheduleServiceImpl(r repository.ScheduleRepository, v *validator.Validate) *ScheduleServiceImpl {
	return &ScheduleServiceImpl{
		r,
		v,
	}
}

var _ ScheduleService = (*ScheduleServiceImpl)(nil)

func (service *ScheduleServiceImpl) All() []entity.Schedule {
	return service.ScheduleRepository.All()
}

func (service *ScheduleServiceImpl) ByJudul(judul string) []entity.Schedule {
	return service.ScheduleRepository.ByJudul(judul)
}
func (service *ScheduleServiceImpl) FindById(id uint64) entity.Schedule {
	return service.ScheduleRepository.FindById(id)
}
func (service *ScheduleServiceImpl) Create(dto dto.ScheduleInsertDTO) (entity.Schedule, error) {
	err := service.Validate.Struct(dto)
	if err != nil {
		return entity.Schedule{}, err
	}
	var Schedule entity.Schedule
	layout := "2006-01-02 15:04:05"
	Schedule.Judul = dto.Judul
	Schedule.StartAt, _ = time.Parse(layout, dto.StartAt)
	Schedule.EndAt, _ = time.Parse(layout, dto.EndAt)
	Schedule.Activities = dto.Activities

	return service.ScheduleRepository.Insert(Schedule), nil
}
func (service *ScheduleServiceImpl) Update(dto dto.ScheduleUpdateDTO) (entity.Schedule, error) {
	err := service.Validate.Struct(dto)
	if err != nil {
		return entity.Schedule{}, err
	}

	var Schedule entity.Schedule
	layout := "2006-01-02 15:04:05"
	Schedule.Id = dto.Id
	Schedule.Judul = dto.Judul
	Schedule.StartAt, _ = time.Parse(layout, dto.StartAt)
	Schedule.EndAt, _ = time.Parse(layout, dto.EndAt)
	return service.ScheduleRepository.Update(Schedule), nil
}
func (service *ScheduleServiceImpl) Delete(id uint64) {
	service.ScheduleRepository.Delete(id)
}
func (service *ScheduleServiceImpl) OnGoingSchedule(from string, to string) []entity.Schedule {
	return service.ScheduleRepository.OnGoingSchedule(from, to)
}

