package handler

import (
	"github.com/gin-gonic/gin"
	"kursachDB/internal/domain/models"
	"kursachDB/internal/services"
)

type Schedule interface {
	Add(schedule services.AddSchedule) error
	Delete(id string) error
	GetAll() []models.Schedule
}

// @Summary AddSchedule
// @Tags schedule
// @Description Add new schedule
// @ID add-schedule
// @Accept json
// @Produce json
// @Router /api/schedule/add [post]
func (h *Handler) AddSchedule(ctx *gin.Context) {
	panic("implement me")
}

// @Summary DeleteSchedule
// @Tags schedule
// @Description Delete schedule
// @ID delete-schedule
// @Accept json
// @Produce json
// @Router /api/schedule/delete [delete]
func (h *Handler) DeleteSchedule(ctx *gin.Context) {
	panic("implement me")
}

// @Summary GetAllSchedule
// @Tags schedule
// @Description Get all schedule
// @ID get-all-schedule
// @Accept json
// @Produce json
// @Router /api/schedule/ [get]
func (h *Handler) GetAllSchedule(ctx *gin.Context) {
	panic("implement me")
}
