package service

import (
	"note_scheduler/dto"
	"note_scheduler/entity"
	"note_scheduler/helper"
	"note_scheduler/repository"

	"github.com/go-playground/validator/v10"
)

type ActivityServiceImpl struct {
	repository.ActivityRepository
	repository.ScheduleRepository
	*validator.Validate
}

func NewActivityServiceImpl(rA repository.ActivityRepository,rS repository.ScheduleRepository, v *validator.Validate) *ActivityServiceImpl {
	return &ActivityServiceImpl{
		rA,
		rS,
		v,
	}
}

var _ ActivityService = (*ActivityServiceImpl)(nil)

func (service *ActivityServiceImpl) All() []entity.Activity {
	return service.ActivityRepository.All()
}

func (service *ActivityServiceImpl) FindById(id uint64) entity.Activity {

	return service.ActivityRepository.FindById(id)
}

func (service *ActivityServiceImpl) Where(e entity.Activity) []entity.Activity {

	return service.ActivityRepository.Where(e)
}

func (service *ActivityServiceImpl) Create(dto dto.ActivityInsertDTO) (entity.Activity, error) {
	err := service.Validate.Struct(dto)
	if err != nil {
		panic(err.Error())
	}
	Activity := entity.Activity{}
	helper.FillStruct(dto, &Activity)
	return service.ActivityRepository.Insert(Activity), nil
}

func (service *ActivityServiceImpl) Update(dto dto.ActivityUpdateDTO) (entity.Schedule, error) {
	err := service.Validate.Struct(dto)
	if err != nil {
		panic(err.Error())
	}
	Activity := entity.Activity{}
	helper.FillStruct(dto, &Activity)
	activity_update := service.ActivityRepository.Update(Activity)
	schedule := service.ScheduleRepository.FindById(activity_update.ScheduleId)
	return schedule, nil
}

func (service *ActivityServiceImpl) Delete(id uint64) entity.Schedule {
	Activity := service.ActivityRepository.FindById(id)
	service.ActivityRepository.Delete(id)
	Schedule := service.ScheduleRepository.FindById(Activity.ScheduleId)
	return Schedule
}
