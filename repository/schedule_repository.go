package repository

import (
	"note_scheduler/entity"

	"gorm.io/gorm"
)

type ScheduleRepositoryImpl struct {
	Conn *gorm.DB
}

func NewScheduleRepositoryImpl(db *gorm.DB) *ScheduleRepositoryImpl {
	return &ScheduleRepositoryImpl{
		Conn: db,
	}
}

var _ ScheduleRepository = (*ScheduleRepositoryImpl)(nil)

func (r *ScheduleRepositoryImpl) All() []entity.Schedule {
	var schedulers []entity.Schedule
	r.Conn.Preload("Activities").Find(&schedulers)
	return schedulers
}
func (r *ScheduleRepositoryImpl) FindById(id uint64) entity.Schedule {
	var Schedule entity.Schedule
	r.Conn.Preload("Activities").Find(&Schedule, id)
	return Schedule
}
func (r *ScheduleRepositoryImpl) Insert(e entity.Schedule) entity.Schedule {
	r.Conn.Create(&e)
	return e
}

func (r *ScheduleRepositoryImpl) Update(e entity.Schedule) entity.Schedule {
	r.Conn.Save(&e)
	r.Conn.Preload("Activities").First(&e)
	return e
}
func (r *ScheduleRepositoryImpl) Delete(id uint64)  {
	r.Conn.Where("schedule_id = ?", id).Delete(&entity.Activity{})
	r.Conn.Delete(&entity.Schedule{}, id)
}

func (r *ScheduleRepositoryImpl) Where(e *entity.Schedule) entity.Schedule {
	var Schedule entity.Schedule
	r.Conn.Where(&e).Find(&Schedule)
	return Schedule
}

func (r *ScheduleRepositoryImpl) ByJudul(judul string) []entity.Schedule {
	var Schedule []entity.Schedule
	r.Conn.Where("judul LIKE ?", judul + "%").Preload("Activities").Find(&Schedule)
	return Schedule
}

func (r *ScheduleRepositoryImpl) OnGoingSchedule(from string, to string) []entity.Schedule {
	var Schedule []entity.Schedule
	r.Conn.Where("start_at >= (?)", from).Where("end_at <= (?)", to).Preload("Activities").Find(&Schedule)
	return Schedule
}
