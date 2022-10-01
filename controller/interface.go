package controller

import "github.com/gin-gonic/gin"

type ScheduleController interface {
	Index(ctx *gin.Context)
	Show(ctx *gin.Context)
	Store(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	OnGoingSchedule(ctx *gin.Context)
}

type ActivityController interface {
	Index(ctx *gin.Context)
	Show(ctx *gin.Context)
	BySchedule(ctx *gin.Context)
	Store(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}
