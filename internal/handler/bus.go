package handler

import (
	"github.com/gin-gonic/gin"
	"kursachDB/internal/domain/models"
	"kursachDB/internal/handler/responses"
	"kursachDB/internal/services"
	"net/http"
)

type Bus interface {
	Add(transport services.AddBus) error
	Update(transport services.AddBus) error
	Delete(stateNumber string) error
	GetAll() []models.Bus
}

// @Summary AddBus
// @Tags bus
// @Description Add bus
// @ID add-bus
// @Accept json
// @Produce json
// @Router /api/bus/add [post]
func (h *Handler) AddBus(ctx *gin.Context) {
	var input services.AddBus
	if err := ctx.ShouldBindJSON(&input); err != nil {
		responses.NewErrorResponse(ctx, http.StatusBadRequest, ErrInvalidArguments)
		return
	}

	if err := h.services.Bus.Add(input); err != nil {
		responses.NewErrorResponse(ctx, http.StatusInternalServerError, ErrInvalidArguments)
		return
	}

	ctx.JSON(http.StatusOK,
		gin.H{
			"status": "success",
		},
	)
}

func validateAddBus(bus services.AddBus) error {
	panic("implement me")
}

// @Summary UpdateBus
// @Tags bus
// @Description Update bus
// @ID update-bus
// @Accept json
// @Produce json
// @Router /api/bus/update [put]
func (h *Handler) UpdateBus(ctx *gin.Context) {
	panic("implement me")
}

// @Summary UpdatePartialBus
// @Tags bus
// @Description Update partial bus
// @ID update-partial-bus
// @Accept json
// @Produce json
// @Router /api/bus/update [patch]
func (h *Handler) UpdatePartialBus(ctx *gin.Context) {
	panic("implement me")
}

// @Summary DeleteBus
// @Tags bus
// @ID delete-bus
// @Accept json
// @Produce json
// @Router /api/bus/delete [delete]
func (h *Handler) DeleteBus(ctx *gin.Context) {
	stateNumber := ctx.Query("stateNumber")
	if err := h.services.Bus.Delete(stateNumber); err != nil {
		responses.NewErrorResponse(ctx, http.StatusInternalServerError, ErrInvalidArguments)
		return
	}

	ctx.JSON(http.StatusOK,
		gin.H{
			"status": "success",
		},
	)
}

// @Summary GetAllBus
// @Tags bus
// @ID get-all-bus
// @Accept json
// @Produce json
// @Router /api/bus/ [get]
func (h *Handler) GetAllBus(ctx *gin.Context) {
	res := h.services.Bus.GetAll()
	ctx.JSON(http.StatusOK, res)
}
