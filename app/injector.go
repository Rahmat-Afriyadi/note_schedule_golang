//go:build wireinject
// +build wireinject

package app

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"note_scheduler/config"
	"note_scheduler/controller"
	"note_scheduler/repository"
	"note_scheduler/service"
)

var ActivitySet = wire.NewSet(
	repository.NewActivityRepositoryImpl,
	wire.Bind(new(repository.ActivityRepository), new(*repository.ActivityRepositoryImpl)),
	service.NewActivityServiceImpl,
	wire.Bind(new(service.ActivityService), new(*service.ActivityServiceImpl)),
	repository.NewScheduleRepositoryImpl,
	wire.Bind(new(repository.ScheduleRepository), new(*repository.ScheduleRepositoryImpl)),
	service.NewScheduleServiceImpl,
	wire.Bind(new(service.ScheduleService), new(*service.ScheduleServiceImpl)),
	controller.NewActivityControllerImpl,
	wire.Bind(new(controller.ActivityController), new(*controller.ActivityControllerImpl)),
)

func InitializedActivityController() controller.ActivityController {
	wire.Build(
		config.SetupDatabaseConnection,
		validator.New,
		ActivitySet,
	)
	return nil
}

var scheduleSet = wire.NewSet(
	repository.NewScheduleRepositoryImpl,
	wire.Bind(new(repository.ScheduleRepository), new(*repository.ScheduleRepositoryImpl)),
	service.NewScheduleServiceImpl,
	wire.Bind(new(service.ScheduleService), new(*service.ScheduleServiceImpl)),
	controller.NewScheduleControllerImpl,
	wire.Bind(new(controller.ScheduleController), new(*controller.ScheduleControllerImpl)),
)

func InitializedScheduleController() controller.ScheduleController {
	wire.Build(
		config.SetupDatabaseConnection,
		validator.New,
		scheduleSet,
	)
	return nil
}
