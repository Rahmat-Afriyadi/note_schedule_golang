package router

import (
	"note_scheduler/controller"

	"github.com/gin-gonic/gin"
)

func NewScheduleRouter(r *gin.Engine, c controller.ScheduleController) {
	// without middleware
	sheduleGroup := r.Group("schedulers")
	{
		sheduleGroup.GET("/", c.Index)
		sheduleGroup.GET("/show/:id", c.Show)
		sheduleGroup.POST("/ongoing", c.OnGoingSchedule)
		sheduleGroup.POST("/", c.Store)
		sheduleGroup.PUT("/update/:id", c.Update)
		sheduleGroup.DELETE("/delete/:id", c.Delete)

	}
}
