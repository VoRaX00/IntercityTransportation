package handler

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"kursachDB/internal/domain/models"
	"kursachDB/internal/handler/responses"
	"kursachDB/internal/services"
	"kursachDB/internal/services/bus"
	"net/http"
	"unicode/utf8"
)

type Bus interface {
	Add(bus services.AddBus) error
	Delete(stateNumber string) error
	GetAll() ([]models.Bus, error)
	Get(stateNumber string) (models.Bus, error)
}

// @Summary AddBus
// @Tags bus
// @Description Add bus
// @ID add-bus
// @Accept json
// @Produce json
// @Param input body services.AddBus true "Bus info to add"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/bus/add [post]
func (h *Handler) AddBus(ctx *gin.Context) {
	var input services.AddBus
	if err := ctx.ShouldBindJSON(&input); err != nil {
		responses.NewErrorResponse(ctx, http.StatusBadRequest, ErrInvalidArguments)
		return
	}

	err := validateAddBus(input)
	if err != nil {
		responses.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err = h.services.Bus.Add(input); err != nil {
		responses.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK,
		gin.H{
			"status": "success",
		},
	)
}

func validateAddBus(bus services.AddBus) error {
	if utf8.RuneCountInString(bus.StateNumber) != 6 {
		return fmt.Errorf("invalid state number")
	}

	for i, val := range bus.StateNumber {
		if (i == 0 || i > 3) && (val < '0' || val > '9') {
			return fmt.Errorf("invalid state number")
		}

		if i > 0 && i < 4 && (val < 'A' || val > 'Z') {
			return fmt.Errorf("invalid state number")
		}
	}
	return nil
}

// @Summary DeleteBus
// @Tags bus
// @ID delete-bus
// @Accept json
// @Produce json
// @Param stateNumber path string true "State number of the bus"
// @Success 200 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/bus/{stateNumber} [delete]
func (h *Handler) DeleteBus(ctx *gin.Context) {
	stateNumber := ctx.Param("stateNumber")

	if err := h.services.Bus.Delete(stateNumber); err != nil {
		if errors.Is(err, bus.ErrBusNotFound) {
			responses.NewErrorResponse(ctx, http.StatusNotFound, ErrRecordNotFound)
		}
		responses.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":      "success",
		"stateNumber": stateNumber,
	})
}

// @Summary GetBus
// @Tags bus
// @ID get-bus
// @Accept json
// @Produce json
// @Param stateNumber path string true "State number of the bus"
// @Success 200 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/bus/{stateNumber} [get]
func (h *Handler) GetBus(ctx *gin.Context) {
	stateNumber := ctx.Param("stateNumber")

	res, err := h.services.Bus.Get(stateNumber)
	if err != nil {
		if errors.Is(err, bus.ErrBusNotFound) {
			responses.NewErrorResponse(ctx, http.StatusNotFound, ErrRecordNotFound)
			return
		}
		responses.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// @Summary GetAllBus
// @Tags bus
// @ID get-all-bus
// @Accept json
// @Produce json
// @Success 200 {array} models.Bus
// @Failure 500 {object} map[string]string
// @Router /api/bus [get]
func (h *Handler) GetAllBus(ctx *gin.Context) {
	buses, err := h.services.Bus.GetAll()
	if err != nil {
		responses.NewErrorResponse(ctx, http.StatusInternalServerError, ErrInvalidArguments)
		return
	}

	ctx.JSON(http.StatusOK, buses)
}
