package main

import (
	"note_scheduler/app"
	"note_scheduler/config"
	"note_scheduler/controller"
	"note_scheduler/router"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

var (
	db *gorm.DB           = config.SetupDatabaseConnection()
	v  validator.Validate = *validator.New()

	// controller
	scheduleController     controller.ScheduleController     = app.InitializedScheduleController()
	activityController controller.ActivityController = app.InitializedActivityController()
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	router.MainRouter(r)

	r.Run(":8080")
}
