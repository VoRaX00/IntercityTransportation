package handler

import (
	"github.com/gin-gonic/gin"
	"kursachDB/internal/domain/models"
	"kursachDB/internal/handler/responses"
	"kursachDB/internal/services"
	"net/http"
)

type Bus interface {
	Add(bus services.AddBus) error
	Delete(stateNumber string) error
	GetAll() ([]models.Bus, error)
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
	buses, err := h.services.Bus.GetAll()
	if err != nil {
		responses.NewErrorResponse(ctx, http.StatusInternalServerError, ErrInvalidArguments)
		return
	}
	ctx.JSON(http.StatusOK, buses)
}
