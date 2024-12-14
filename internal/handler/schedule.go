package handler

import (
	"github.com/gin-gonic/gin"
	"kursachDB/internal/domain/models"
	"kursachDB/internal/handler/responses"
	"kursachDB/internal/services"
	"net/http"
	"strconv"
)

type Schedule interface {
	Add(schedule services.AddSchedule) error
	Delete(id int) error
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
	var input services.AddSchedule
	if err := ctx.ShouldBindJSON(&input); err != nil {
		responses.NewErrorResponse(ctx, http.StatusBadRequest, ErrInvalidArguments)
		return
	}

	if err := validateAddSchedule(input); err != nil {
		responses.NewErrorResponse(ctx, http.StatusBadRequest, ErrInvalidArguments)
		return
	}

	if err := h.services.Schedule.Add(input); err != nil {
		responses.NewErrorResponse(ctx, http.StatusInternalServerError, ErrInternalServer)
		return
	}

	ctx.JSON(http.StatusOK,
		gin.H{
			"status": "success",
		},
	)
}

func validateAddSchedule(schedule services.AddSchedule) error {
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
	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		responses.NewErrorResponse(ctx, http.StatusBadRequest, ErrInvalidArguments)
		return
	}

	if err = h.services.Schedule.Delete(id); err != nil {
		responses.NewErrorResponse(ctx, http.StatusInternalServerError, ErrInternalServer)
		return
	}

	ctx.JSON(http.StatusOK,
		gin.H{
			"status": "success",
		},
	)
}

// @Summary GetAllSchedule
// @Tags schedule
// @Description Get all schedule
// @ID get-all-schedule
// @Accept json
// @Produce json
// @Router /api/schedule/ [get]
func (h *Handler) GetAllSchedule(ctx *gin.Context) {
	res := h.services.Schedule.GetAll()
	ctx.JSON(http.StatusOK, res)
}
