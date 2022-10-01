package entity


type Activity struct {
	Id       uint64 `gorm:"primaryKey:autoIncreament" json:"id"`
	Activity string `json:"activity"`

	ScheduleId uint64   `gorm:"not null" json:"schedule_id"`
	Schedule   Schedule `gorm:"foreign:ScheduleId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}
