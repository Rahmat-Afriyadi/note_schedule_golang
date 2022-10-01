package controller

import (
	"net/http"
	"note_scheduler/dto"
	"note_scheduler/entity"
	"note_scheduler/helper"
	"note_scheduler/service"

	// "os/user"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ScheduleControllerImpl struct {
	service.ScheduleService
}

func NewScheduleControllerImpl(s service.ScheduleService) *ScheduleControllerImpl {
	return &ScheduleControllerImpl{
		s,
	}
}

var _ ScheduleController = (*ScheduleControllerImpl)(nil)

func (controller *ScheduleControllerImpl) Index(ctx *gin.Context) {
	var schedules []entity.Schedule
	if ctx.Query("judul") != "" {
		schedules = controller.ScheduleService.ByJudul(ctx.Query("judul"))	
	} else {
		schedules = controller.ScheduleService.All()
	}
	response := helper.BuildResponse(200,schedules)
	ctx.JSON(http.StatusAccepted, response)
	return
}
func (controller *ScheduleControllerImpl) OnGoingSchedule(ctx *gin.Context) {
	dto := dto.OnGoingSchedule{}
	err := ctx.ShouldBind(&dto)
	if err != nil {
		res := helper.BuildErrorResponse(403, err.Error())
		ctx.JSON(http.StatusBadRequest, res)
	}
	result := controller.ScheduleService.OnGoingSchedule(dto.From, dto.To)
	response := helper.BuildResponse(200,  result)
	ctx.JSON(http.StatusAccepted, response)
	return
}
func (controller *ScheduleControllerImpl) Show(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	result := controller.ScheduleService.FindById(id)
	response := helper.BuildResponse(200, result)
	ctx.JSON(http.StatusAccepted, response)
	return
}
func (controller *ScheduleControllerImpl) Store(ctx *gin.Context) {
	dto := dto.ScheduleInsertDTO{}
	err := ctx.ShouldBindJSON(&dto)
	if err != nil {
		res := helper.BuildErrorResponse(403, err.Error())
		ctx.JSON(http.StatusBadRequest, res)
	}
	result, err := controller.ScheduleService.Create(dto)
	if err != nil {
		res := helper.BuildErrorResponse(403, err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	response := helper.BuildResponse(200, result)
	ctx.JSON(http.StatusAccepted, response)
	return
}

func (controller *ScheduleControllerImpl) Update(ctx *gin.Context) {
	dto := dto.ScheduleUpdateDTO{}
	err := ctx.ShouldBind(&dto)
	if err != nil {
		panic(err.Error())	
	}
	dto.Id, _ = strconv.ParseUint(ctx.Param("id"), 10, 64)
	result, err := controller.ScheduleService.Update(dto)
	if err != nil {
		response := helper.BuildErrorResponse(403, err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.BuildResponse(200, result)
	ctx.JSON(http.StatusAccepted, response)
	return
}
func (controller *ScheduleControllerImpl) Delete(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	schedule := controller.ScheduleService.FindById(id)
	controller.ScheduleService.Delete(id)
	res := helper.BuildResponse(200, schedule)
	ctx.JSON(http.StatusAccepted, res)
	return
}
