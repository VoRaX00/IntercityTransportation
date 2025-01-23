package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"kursachDB/internal/domain/models"
	"kursachDB/internal/handler/responses"
	"kursachDB/internal/services"
	"net/http"
	"time"
)

type Flight interface {
	Add(flight services.AddFlight) error
	GetAll() ([]models.Flight, error)
}

// @Summary AddFlight
// @Tags flight
// @Description Add new flight
// @ID add-flight
// @Accept json
// @Produce json
// @Param input body services.AddFlight true "Flight info for add"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/flight/add [post]
func (h *Handler) AddFlight(ctx *gin.Context) {
	var input services.AddFlight
	if err := ctx.ShouldBindJSON(&input); err != nil {
		responses.NewErrorResponse(ctx, http.StatusBadRequest, ErrInvalidArguments)
		return
	}

	if err := validateAddFlight(input); err != nil {
		responses.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Flight.Add(input); err != nil {
		responses.NewErrorResponse(ctx, http.StatusInternalServerError, ErrInternalServer)
		return
	}

	ctx.JSON(http.StatusOK,
		gin.H{
			"status": "success",
		},
	)
}

func validateAddFlight(flight services.AddFlight) error {
	departure, err := time.Parse("02.01.2006 15:04", flight.Departure)
	if err != nil {
		return fmt.Errorf("departure date format is wrong")
	}

	arrival, err := time.Parse("02.01.2006 15:04", flight.Arrival)
	if err != nil {
		return fmt.Errorf("arrival date format is wrong")
	}

	if departure.After(arrival) {
		return fmt.Errorf("departure date is wrong")
	}
	return nil
}

// @Summary GetAllFlight
// @Tags flight
// @Description Get all flight
// @ID get-all-flight
// @Accept json
// @Produce json
// @Success 200 {array} models.Flight
// @Failure 500 {object} map[string]string
// @Router /api/flight [get]
func (h *Handler) GetAllFlight(ctx *gin.Context) {
	res, err := h.services.Flight.GetAll()
	if err != nil {
		responses.NewErrorResponse(ctx, http.StatusInternalServerError, ErrInternalServer)
		return
	}
	ctx.JSON(http.StatusOK, res)
}
