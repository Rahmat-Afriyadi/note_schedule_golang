package router

import (
	"note_scheduler/controller"

	"github.com/gin-gonic/gin"
)

func NewActivityRouter(r *gin.Engine, c controller.ActivityController) {
	activityGroup := r.Group("activities")
	{
		activityGroup.GET("/", c.Index)
		activityGroup.GET("/show/:id", c.Show)
		activityGroup.GET("/by-shedule/:schedule_id", c.BySchedule)
		activityGroup.POST("/store/:schedule_id", c.Store)
		activityGroup.PUT("/update/:id", c.Update)
		activityGroup.DELETE("/delete/:id", c.Delete)

	}
}
