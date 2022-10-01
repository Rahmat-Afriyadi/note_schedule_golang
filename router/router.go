package router

import (
	"note_scheduler/app"
	"note_scheduler/controller"

	"github.com/gin-gonic/gin"
)

var (
	scheduleController     controller.ScheduleController     = app.InitializedScheduleController()
	activityController controller.ActivityController = app.InitializedActivityController()
)

func MainRouter(r *gin.Engine) {
	NewScheduleRouter(r, scheduleController)
	NewActivityRouter(r, activityController)
}
