package controller

import (
	"net/http"
	"note_scheduler/dto"
	"note_scheduler/entity"
	"note_scheduler/helper"
	"note_scheduler/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ActivityControllerImpl struct {
	service.ActivityService
	service.ScheduleService
}

func NewActivityControllerImpl(s service.ActivityService, ss service.ScheduleService) *ActivityControllerImpl {
	return &ActivityControllerImpl{
		s,
		ss,
	}
}

func (controller *ActivityControllerImpl) Index(ctx *gin.Context) {
	res := controller.ActivityService.All()
	helper.BuildResponse(200, res)
	return
}

func (controller *ActivityControllerImpl) Show(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	res := controller.ActivityService.FindById(id)	
	response := helper.BuildResponse(200, res)
	ctx.JSON(http.StatusAccepted, response)
	return
}
func (controller *ActivityControllerImpl) BySchedule(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("schedule_id"), 10, 64)
	activity := entity.Activity{
		ScheduleId: id,
	}
	res := controller.ActivityService.Where(activity)
	
	response := helper.BuildResponse(200,  res)
	ctx.JSON(http.StatusAccepted, response)
	return
}

func (controller *ActivityControllerImpl) Store(ctx *gin.Context) {
	dto := dto.ActivityInsertDTO{}
	err := ctx.ShouldBind(&dto)
	if err != nil {
		res := helper.BuildErrorResponse(403, err.Error())
		ctx.JSON(http.StatusBadRequest, res)
	}
	dto.ScheduleId, _ = strconv.ParseUint(ctx.Param("schedule_id"), 10, 64)
	result, err := controller.ActivityService.Create(dto)
	if err != nil {
		res := helper.BuildErrorResponse(403, err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	schedule := controller.ScheduleService.FindById(result.ScheduleId)
	response := helper.BuildResponse(200,  schedule)
	ctx.JSON(http.StatusAccepted, response)
	return

}
func (controller *ActivityControllerImpl) Update(ctx *gin.Context) {
	dto := dto.ActivityUpdateDTO{}
	err := ctx.ShouldBind(&dto)
	if err != nil {
		res := helper.BuildErrorResponse(403, err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	dto.Id, _ =strconv.ParseUint(ctx.Param("id"), 10, 64) 
	result, err := controller.ActivityService.Update(dto)
	if err != nil {
		res := helper.BuildErrorResponse(403, err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	response := helper.BuildResponse(200,  result)
	ctx.JSON(http.StatusAccepted, response)
	return
}
func (controller *ActivityControllerImpl) Delete(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	schedule := controller.ActivityService.Delete(id)
	res := helper.BuildResponse(200,  schedule)
	ctx.JSON(http.StatusAccepted, res)
	return
}
