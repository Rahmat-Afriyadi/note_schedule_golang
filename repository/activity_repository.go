package repository

import (
	"note_scheduler/entity"

	"gorm.io/gorm"
)

type ActivityRepositoryImpl struct {
	Conn *gorm.DB
}

func NewActivityRepositoryImpl(db *gorm.DB) *ActivityRepositoryImpl {
	return &ActivityRepositoryImpl{
		Conn: db,
	}
}

var _ ActivityRepository = (*ActivityRepositoryImpl)(nil)

func (r *ActivityRepositoryImpl) All() []entity.Activity {
	var Activitys []entity.Activity
	r.Conn.Find(&Activitys)
	return Activitys
}
func (r *ActivityRepositoryImpl) FindById(id uint64) entity.Activity {
	var Activity entity.Activity
	r.Conn.Find(&Activity, id)
	return Activity
}
func (r *ActivityRepositoryImpl) Insert(e entity.Activity) entity.Activity {
	r.Conn.Save(&e)
	return e
}

func (r *ActivityRepositoryImpl) Update(e entity.Activity) entity.Activity {
	r.Conn.Save(&e)
	return e
}
func (r *ActivityRepositoryImpl) Delete(id uint64) error {
	r.Conn.Delete(&entity.Activity{},id)
	return nil
}
func (r *ActivityRepositoryImpl) Where(e entity.Activity) []entity.Activity {
	var Activity []entity.Activity
	r.Conn.Where(&e).Find(&Activity)
	return Activity
}
