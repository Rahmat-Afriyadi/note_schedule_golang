package dto

type ActivityInsertDTO struct {
	Activity   string `json:"activity" validate:"required" form:"activity" bind:"required"`
	ScheduleId uint64 `json:"schedule_id" validate:"required"`
}

type ActivityUpdateDTO struct {
	Id         uint64 `json:"id" validate:"required" form:"id" bind:"required"`
	Activity   string `json:"activity" validate:"required" form:"activity" bind:"required"`
	ScheduleId uint64 `json:"schedule_id" validate:"required" form:"schedule_id" bind:"required"`
}
